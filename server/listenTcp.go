package server

import (
	"encoding/json"
	"gin-go/model"
	"gin-go/service"
	"log"
	"net"
	"reflect"

	"github.com/patrickmn/go-cache"
)

func ServerTCP(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Printf("接受客户端连接异常: %+v", err.Error())
			continue
		}
		log.Printf("数据上报客户端连接来自: %+v", conn.RemoteAddr().String())
		// defer conn.Close()
		//init device

		go func() {
			defer func() {
				log.Print("tcp exit")
				conn.Close()
			}()

			data := make([]byte, 10240)
			for {
				//解析TCP里头的Json信息
				i, err := conn.Read(data)
				tcpjson := model.DeviceReportInfo{}
				t := string(data[0:i])

				// log.Printf(t)
				if err := json.Unmarshal([]byte(t), &tcpjson); err != nil {
					// 如果解析过程中出现错误，打印错误信息并退出
					log.Printf("Error: %+v", err)
					break
				}

				// json.Unmarshal(data[0:i], &tcpjson)
				tcpjson.DeviceIp = conn.RemoteAddr().String()
				if tcpjson.DeviceName == "" {
					conn.Write([]byte{'e', 'r', 'r'})
					continue
				} else {
					// tmpjson.DeviceState = true
					val, found := service.DevCache.Get(tcpjson.DeviceName)

					if found {
						// 定义一个 Person 类型的变量
						p := model.DeviceReportInfo{}

						// 判断 val 是否为 DeviceReportInfo 类型，并将它赋值给 p
						if v, ok := val.(model.DeviceReportInfo); ok {
							p = v
						}
						//对比上报的设备和在内存内的设备信息是否相等
						if reflect.DeepEqual(p, tcpjson) {
							// p1 和 p2 相等
							//log.Printf("相等")
							//log.Printf("Decive %s is in cache!Not thing to do", tcpjson.DeviceName)
						} else {
							log.Printf("Decive %s is in cache,but not Equal", tcpjson.DeviceName)
							service.DevCache.Set(tcpjson.DeviceName, tcpjson, cache.DefaultExpiration)
							// log.Printf(tcpjson)
							log.Printf("Error occurred: %+v", tcpjson)
							//log.Printf("不相等")
							//通知所有web客户端
							for _, val := range service.ChanMap {
								// 检查通道是否已经关闭
								if val == nil {
									// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
									//delete(service.ChanMap, key)
									continue
								}

								// 发送消息

								select {
								case val <- model.WebSocket{
									Data: service.DevCache.Items(),
									Cmd:  1,
								}:
									log.Printf("send data to client")
									// 如果通道没有关闭，则可以向通道发送数据
								case <-val:
									log.Printf("Client failed to send")
									// 如果通道已经关闭，则执行特定的操作
								}
							}
						}

						// service.DevCache.()
					} else {
						log.Printf("Decive %s is Not in cache,New Device!", tcpjson.DeviceName)
						log.Printf("data:%+v", tcpjson)
						service.DevCache.Add(tcpjson.DeviceName, tcpjson, cache.DefaultExpiration)
						//通知所有web客户端
						for _, val := range service.ChanMap {
							// 检查通道是否已经关闭
							if val == nil {
								// 如果通道已经关闭，可以执行相应的处理（例如删除键值对）
								//delete(service.ChanMap, key)
								continue
							}

							// 发送消息

							select {
							case val <- model.WebSocket{
								Data: service.DevCache.Items(),
								Cmd:  1,
							}:
								log.Printf("send data to client")
								// 如果通道没有关闭，则可以向通道发送数据
							case <-val:
								log.Printf("Client failed to send")
								// 如果通道已经关闭，则执行特定的操作
							}
						}
					}
				}

				//log.Printf("andao send")
				if err != nil {
					log.Printf("读取客户端数据错误: %+v", err.Error())
					break
				}

				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}
		}()
	}
}

func ServerControlTCP(listen *net.TCPListener) {
	for {
		// 接受客户端连接

		conn, err := listen.AcceptTCP()
		if err != nil {
			// 处理错误
			log.Printf("接受客户端连接异常: %+v", err.Error())
			continue
		}

		// 获取客户端地址
		addr := conn.RemoteAddr().String()

		// 输出连接的客户端地址
		log.Printf("控制硬件客户端连接来自: %+v", addr)

		// 创建一个通道
		ch := make(chan []byte)
		out := make(chan string)

		// 将客户端地址和通道绑定
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			// 处理错误
			log.Printf("Error occurred: %+v", err)
		}
		service.DevClients[host] = ch
		service.DevClientsOut[host] = out

		conn.Write([]byte("start listen"))
		go func() {
			for {
				// 接收字符串
				b, ok := <-ch
				if !ok {
					// 通道已经关闭，退出循环
					break
				}
				log.Printf("Received from channel: %+v", b)
				// 写入客户端
				conn.Write(b)
			}
		}()
		// 开启一个新的 goroutine 处理连接
		go func() {
			defer func() {
				log.Print("tcp exit")
				conn.Close()
				close(ch)
				close(out)
				delete(service.DevClients, addr)
			}()
			// 在函数退出时关闭连接
			for {

				// 启动一个 goroutine 等待其他 goroutine 发送字符串

				// 创建一个字节数组来存储数据
				data := make([]byte, 10240)

				// 读取数据
				i, err := conn.Read(data)
				if err != nil {
					log.Printf("Error: %+v", err)
					// 处理错误
					break
				}

				// 将读取到的字节数组转换为字符串
				t := string(data[0:i])
				tcpjson := model.DeviceSerchInfo{}

				// log.Printf(t)
				if err := json.Unmarshal([]byte(t), &tcpjson); err != nil {
					// 如果解析过程中出现错误，打印错误信息并退出
					log.Printf("Error: %+v", err)
					// return
				}
				out <- t

				// tcpjson.GoProIP = "172.25.149.51"
				// tcpjson.DeviceSubIp = "172.25.149.52"
				// tcpjson.VideoPort = 12001
				// json_p, err := json.Marshal(tcpjson)
				// if err != nil {
				// 	log.Printf("转换json字符串失败！")
				// }
				// cmd := []byte{0xAA, 0x55, 0x00, 0xff, 0xF0, 0x00, 0x00, 0x00}
				// conn.Write(append(cmd, json_p...))

				// conn.Write([]byte{0xAA, 0x55, 0x00, 0x00, 0xF4, 0x00, 0x00, 0x00})
				// log.Printf("Test:", t)
			}
		}()
	}
}
func ListenTCP() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 11119, Zone: ""})
	if err != nil {
		log.Printf("监听端口失败: %+v", err.Error())
		return
	}
	log.Printf("已初始化连接，等待客户端连接...")
	go ServerTCP(listen)
	// fmt.Print("已初始化end")
}

func ListenControlTCP() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 11118, Zone: ""})
	if err != nil {
		log.Printf("监听端口失败: %+v", err.Error())
		return
	}
	log.Printf("已初始化连接，等待客户端连接...")
	go ServerControlTCP(listen)
	// fmt.Print("已初始化end")
}
