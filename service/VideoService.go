package service

import (
	"fmt"
	"gin-go/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type fileType struct {
	Name    string    `json:"Name"`
	Size    int       `json:"Size"`
	ModTime time.Time `json:"ModTime"`
}

// 用来取消视频录制任务的
var TimedVideoTaskCancelVar = false

func GetLocalVideoFileList(c *gin.Context) {
	files, _ := ioutil.ReadDir("../VideoDownload")
	var arr []fileType
	for _, f := range files {
		arr = append(arr, fileType{
			Name:    f.Name(),
			Size:    int(f.Size()),
			ModTime: f.ModTime(),
		})
		// log.Printf(f.Name(), f.IsDir(), f.Size(), f.ModTime())
	}
	c.JSONP(http.StatusOK, arr)
}

func OpenFolder(c *gin.Context) {
	// 先返回上一级目录
	dirname := "./VideoDownload"
	parent, err := os.Getwd()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    777,
			"message": err,
		})
	}
	separator := string(os.PathSeparator)
	parentIndex := strings.LastIndex(parent, separator)
	parent = parent[:parentIndex]

	// 根据指定的文件夹进入打开
	var path string
	if dirname[0] == '.' && dirname[1] == '/' {
		relativePath := dirname[2:]
		relativePath = strings.ReplaceAll(relativePath, "/", "\\") // 转换为 windows 路径格式
		path = parent + separator + relativePath
	} else {
		path = dirname
	}

	// 打开文件夹并前置
	cmd := exec.Command("cmd", "/c", "start", "explorer", "select,", path)
	err = cmd.Start()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    777,
			"message": err,
		})
	}

}

func RenameLocalVideoFile(c *gin.Context) {
	// 重命名文件
	oldName := c.PostForm("oldName")
	newName := c.PostForm("newName")

	file := "../VideoDownload/" + oldName
	err1 := os.Rename(file, "../VideoDownload/"+newName)
	if err1 != nil {
		c.JSON(200, gin.H{
			"code":    777,
			"message": err1,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    666,
			"message": "文件重命名成功",
		})
	}
}

func RemoveLocalVideoFile(c *gin.Context) {
	name := c.PostForm("name")
	err1 := os.Remove("../VideoDownload/" + name)
	if err1 != nil {
		c.JSON(200, gin.H{
			"code":    777,
			"message": err1,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    666,
			"message": "file remove OK!",
		})
	}
}

func DownloadFile(c *gin.Context) {
	// 获取 URL 和文件名
	url := c.Query("url")
	fileName := c.Query("fileName")

	// 创建文件
	file, err := os.Create("../VideoDownload/" + fileName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer file.Close()

	// 创建 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()

	// 创建进度条
	progress := &ProgressTracker{FileName: fileName, Total: resp.ContentLength}
	reader := io.TeeReader(resp.Body, progress)

	// 获取响应关闭通知通道
	closeNotify := c.Writer.CloseNotify()

	go func() {
		var repeatTimes int = 0
		var tmpProgress int64
		for {
			// 这里主要防止重复
			if tmpProgress == progress.Progress {
				repeatTimes += 1
				if repeatTimes > 30 {
					log.Printf("repeatTimes!")
					resp.Body.Close()

					file.Close()

					// 告诉前端下载失败,并且删除视频文件。
					err := os.Remove("../VideoDownload/" + fileName)
					time.Sleep(time.Millisecond * 1300)
					if err != nil {
						fmt.Println("删除文件失败:", err)
						return
					}
					for _, val := range ChanMap {
						progress.Done = true
						// 检查通道是否已经关闭
						if val == nil {
							// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
							// delete(service.ChanMap, key)
							log.Printf("channel is closed!")
							continue
						}
						// 发送消息

						select {
						case val <- model.WebSocket{
							Data: "视频文件下载失败：" + fileName,
							Cmd:  4,
						}:
							// log.Print("test")
							// log.Print(progress)
							log.Printf("send data to client  -end")
							// 如果通道没有关闭，则可以向通道发送数据
						case <-val:
							log.Printf("Client failed to send -end")
							// 如果通道已经关闭，则执行特定的操作
						}
					}
					break
				}
			}

			tmpProgress = progress.Progress
			time.Sleep(time.Millisecond * 1300)

			for _, val := range ChanMap {
				progress.Done = false
				// 检查通道是否已经关闭
				if val == nil {
					// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
					//delete(service.ChanMap, key)
					continue
				}

				// 发送消息
				select {
				case val <- model.WebSocket{
					Data: progress,
					Cmd:  2,
				}:
					log.Printf("send data to client -1")
					// 如果通道没有关闭，则可以向通道发送数据
				case <-val:
					log.Printf("Client failed to send -1")
					// 如果通道已经关闭，则执行特定的操作
				case <-closeNotify:
					// 响应已经关闭，执行相应的处理
					log.Println("Response closed")
				}
			}

			if progress.Progress > 0 && progress.Total == progress.Progress {
				// log.Print(progress)
				for _, val := range ChanMap {
					progress.Done = true
					// 检查通道是否已经关闭
					if val == nil {
						// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
						//delete(service.ChanMap, key)
						log.Printf("channel is closed!")
						continue
					}
					// 发送消息
					for {
						select {
						case val <- model.WebSocket{
							Data: progress,
							Cmd:  2,
						}:
							// log.Print("test")
							// log.Print(progress)
							log.Printf("send data to client")
							return
							// 如果通道没有关闭，则可以向通道发送数据
						case <-val:
							log.Printf("Client failed to send -end")
							// 如果通道已经关闭，则执行特定的操作
						}
					}
				}
				if progress.Progress < 0 {
					log.Printf("Progress Number is -1")
					break
				}
				break
			}
		}
	}()
	// 使用 io.Copy() 函数将文件内容写入磁盘
	io.Copy(file, reader)

	// 返回下载完成
	c.String(200, "download completed")
}

// ProgressTracker 结构体实现了 io.Writer 接口，并跟踪进度
type ProgressTracker struct {
	FileName string `json:"file_name"` // 下载文件的名称
	Total    int64  `json:"total"`     // 下载文件的总字节数
	Progress int64  `json:"progress"`  // 下载进度
	Done     bool   `json:"done"`
}

// Write 实现了 io.Writer 接口的 Write() 方法，用于更新进度条
func (p *ProgressTracker) Write(b []byte) (int, error) {
	n := len(b)
	// p.total += int64(n)
	p.Progress += int64(n)

	// 在这里处理文件名
	//log.Printf("File name:", p.fileName)
	//通知所有web客户端

	// 在这里更新进度条
	// log.Printf("File name:%s Progress: %d/%d bytes\n", p.FileName, p.Progress, p.Total)
	return n, nil
}

// 时间任务结构体
type TimedVideoTaskStatus struct {
	StartTime int64 `json:"start_time"` //开始录制视频时间
	EndTime   int64 `json:"end_time"`   //结束录制视频时间
}

func TimedVideoTaskCancel(c *gin.Context) {
	TimedVideoTaskCancelVar = true
	log.Printf("TimedVideoTaskCancel")

}

func TimedVideoTask(c *gin.Context) {
	// 获取开始时间戳
	startTimeStr := c.Query("startTime")
	// 获取结束时间戳
	endTimeStr := c.Query("endTime")
	// 创建一个定时器，每分钟检查一次当前时间
	timer := time.NewTicker(time.Second)

	// 先将字符串转换为整数
	endTimeStamp, err := strconv.ParseInt(endTimeStr, 10, 64)
	if err != nil {
		// 如果转换失败，处理错误
		// ...
	}
	endTime := time.Unix(endTimeStamp, 0)

	// 先将字符串转换为整数
	startTimeStamp, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		// 如果转换失败，处理错误
		// ...
	}
	startTime := time.Unix(startTimeStamp, 0)

	tvts := TimedVideoTaskStatus{
		StartTime: startTimeStamp * 1000,
		EndTime:   endTimeStamp * 1000,
	}

	TimedVideoTaskCancelVar = false

	go func() {
		defer func() {
			for _, val := range ChanMap {
				// 检查通道是否已经关闭
				if val == nil {
					// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
					//delete(service.ChanMap, key)
					continue
				}

				// 发送消息
				select {
				case val <- model.WebSocket{
					Data: nil, //结束
					Cmd:  3,
				}:
					log.Printf("send data to client")
					// 如果通道没有关闭，则可以向通道发送数据
				case <-val:
					log.Printf("Client failed to send")
					// 如果通道已经关闭，则执行特定的操作
				}
			}
		}()
		for {
			if TimedVideoTaskCancelVar {
				break
			}

			// 获取当前时间
			now := <-timer.C
			log.Printf("等待开始时间:" + startTime.Format("2006-01-02 15:04:05"))
			// 如果当前时间在开始录像时间之后，且在结束录像时间之前，则开始录像

			if now.After(startTime) && now.Before(endTime) {
				log.Printf("Start recording video")

				for _, v := range DevCache.Items() {
					go func(v cache.Item) {
						for i := 0; i < 3; i++ {
							fmt.Println(v.Object.(model.DeviceReportInfo).GoProIP)
							ip := v.Object.(model.DeviceReportInfo).GoProIP
							client := &http.Client{}
							//设置usb可以控制
							// 发送第一个HTTP GET请求
							resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=1")
							if err != nil {
								for _, val := range ChanMap {
									// progress.Done = true
									// 检查通道是否已经关闭
									if val == nil {
										// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
										//delete(service.ChanMap, key)
										log.Printf("channel is closed!")
										continue
									}
									// 发送消息

									select {
									case val <- model.WebSocket{
										Data: fmt.Sprintf("相机:%s,开始定时录制失败,重试第%v次", ip, i+1),
										Cmd:  4,
									}:
										// log.Print("test")
										// log.Print(progress)
										log.Printf("send data to client  -end")
										// 如果通道没有关闭，则可以向通道发送数据
									case <-val:
										log.Printf("Client failed to send -end")
										// 如果通道已经关闭，则执行特定的操作
									}
								}
								log.Printf("resp: %+v", resp)
								log.Printf("Error occurred: %+v", err)
								continue
							}
							//设置usb设置为录制视频的窗口
							// 发送第二个HTTP GET请求
							resp, err = client.Get("http://" + ip + ":8080/gopro/camera/presets/set_group?id=1000")
							if err != nil {
								log.Printf("resp: %+v", resp)
								log.Printf("Error occurred: %+v", err)
								continue
							}
							//设置usb开启录制视频
							// 发送第三个HTTP GET请求
							resp, err = client.Get("http://" + ip + ":8080/gopro/camera/shutter/start")
							if err != nil {
								log.Printf("Error occurred: %+v", err)

								continue
							}
							break
						}
					}(v)

				}
				break
			}
			for _, val := range ChanMap {
				// 检查通道是否已经关闭
				if val == nil {
					// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
					//delete(service.ChanMap, key)
					continue
				}

				// 发送消息
				select {
				case val <- model.WebSocket{
					Data: tvts, //代表视频等待播放，可以
					Cmd:  3,
				}:
					log.Printf("send data to client")
					// 如果通道没有关闭，则可以向通道发送数据
				case <-val:
					log.Printf("Client failed to send")
					// 如果通道已经关闭，代表视频录制中
				}
			}
		}

		for {
			// if TimedVideoTaskCancelVar {
			// 	break
			// }

			// 获取当前时间
			now := <-timer.C
			log.Printf("等待结束时间:" + endTime.Format("2006-01-02 15:04:05"))

			// 如果当前时间在结束录像时间之后，则结束录像
			if now.After(endTime) || TimedVideoTaskCancelVar {
				log.Printf("Stop recording video")
				for _, v := range DevCache.Items() {
					go func(v cache.Item) {
						for i := 0; i < 3; i++ {
							fmt.Println(v.Object.(model.DeviceReportInfo).GoProIP)
							ip := v.Object.(model.DeviceReportInfo).GoProIP
							client := &http.Client{}
							//设置usb可以控制
							// 发送第一个HTTP GET请求
							resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=1")
							if err != nil {
								for _, val := range ChanMap {
									// progress.Done = true
									// 检查通道是否已经关闭
									if val == nil {
										// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
										//delete(service.ChanMap, key)
										log.Printf("channel is closed!")
										continue
									}
									// 发送消息

									select {
									case val <- model.WebSocket{
										Data: fmt.Sprintf("相机:%s,结束定时录制失败,重试第%v次", ip, i+1),
										Cmd:  4,
									}:
										// log.Print("test")
										// log.Print(progress)
										log.Printf("send data to client  -end")
										// 如果通道没有关闭，则可以向通道发送数据
									case <-val:
										log.Printf("Client failed to send -end")
										// 如果通道已经关闭，则执行特定的操作
									}
								}
								log.Printf("Error occurred: %+v", resp)
								log.Printf("Error occurred: %+v", err)
								continue
							}
							//设置usb设置为录制视频的窗口
							// 发送第二个HTTP GET请求
							resp, err = client.Get("http://" + ip + ":8080/gopro/camera/presets/set_group?id=1000")
							if err != nil {
								log.Printf("Error occurred: %+v", resp)
								log.Printf("Error occurred: %+v", err)
								continue
							}
							//设置usb开启录制视频
							// 发送第三个HTTP GET请求
							resp, err = client.Get("http://" + ip + ":8080/gopro/camera/shutter/stop")
							if err != nil {
								log.Printf("Error occurred: %+v", err)

								continue
							}
							break
						}
					}(v)

				}
				break
			}

			for _, val := range ChanMap {
				// 检查通道是否已经关闭
				if val == nil {
					// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
					//delete(service.ChanMap, key)
					continue
				}

				// 发送消息
				select {
				case val <- model.WebSocket{
					Data: tvts, //代表视频录制中
					Cmd:  3,
				}:
					log.Printf("send data to client")
					// 如果通道没有关闭，则可以向通道发送数据
				case <-val:
					log.Printf("Client failed to send")
					// 如果通道已经关闭，则执行特定的操作
				}
			}
		}
	}()

}
