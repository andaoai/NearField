package service

import (
	"encoding/json"
	"fmt"
	"gin-go/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// 存放已有的设备 创建一个默认过期时间为 1 分钟的缓存，每 1 分钟清除一次过期项目
var DevCache = cache.New(1*time.Minute, 1*time.Minute)

var DevClients = make(map[string]chan []byte)
var DevClientsOut = make(map[string]chan string)

func DelDevice(c *gin.Context) {
	DevName := c.Query("device_name")
	DevCache.Delete(DevName)
	for _, val := range ChanMap {
		val <- model.WebSocket{
			Data: DevCache.Items(),
			Cmd:  1,
		}
	}
}
func SearchDeviceInfo(c *gin.Context) {
	addr := c.Query("device_ip")
	// 从 map 中获取客户端对应的通道
	ch := DevClients[addr]
	out := DevClientsOut[addr]
	// 将字符串发送到通道
	ch <- []byte{0xAA, 0x55, 0x00, 0xff, 0xF1, 0x00, 0x00, 0x00}
	// c.JSONP(http.StatusOK, "TEST SEND")

	// 创建一个超时通道，1 秒后向它发送值
	timeout := time.After(time.Second)

	select {
	case s, ok := <-out:
		// 从 out 通道中读取到值
		if ok {
			// 处理 s 的值
			log.Printf(s)
			tmpjson := model.DeviceSerchInfo{}

			// log.Printf(t)
			if err := json.Unmarshal([]byte(s), &tmpjson); err != nil {
				// 如果解析过程中出现错误，打印错误信息并退出
				log.Printf("Error occurred: %+v", err)
				// return
			}
			c.JSONP(http.StatusOK, tmpjson)

		} else {
			// 通道已经关闭
			log.Printf("通道已经关闭")
		}
	case <-timeout:
		c.JSONP(http.StatusOK, "超时通道")
	}

}

func postFormToInt(c *gin.Context, key string) (int, error) {
	// 从表单中获取值
	value := c.PostForm(key)
	// 将字符串转换为整数
	return strconv.Atoi(value)
}

func SetDevice(c *gin.Context) {
	httpjson := &model.DeviceSerchInfo{}
	httpjson.DeviceName = c.PostForm("device_name")
	httpjson.DeviceMainIp = c.PostForm("device_main_ip")
	httpjson.DeviceSubIp = c.PostForm("device_sub_ip")
	httpjson.DeviceGatWay = c.PostForm("device_gateway")
	httpjson.GoProIP = c.PostForm("gopro_ip")

	// 获取并转换值
	videoPort, err := strconv.ParseInt(c.PostForm("video_port"), 10, 64)
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.VideoPort = videoPort

	// 获取并转换值
	baudrate, err := strconv.ParseInt(c.PostForm("baudrate"), 10, 64)
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.Baudrate = baudrate

	// 获取并转换值
	rfModel, err := postFormToInt(c, "rf_mode")
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.RfModel = int8(rfModel)

	// 获取并转换值
	rfBandwidth, err := postFormToInt(c, "rf_bandwidth")
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.RfBandwidth = int8(rfBandwidth)

	// 获取并转换值
	rfKey, err := strconv.ParseInt(c.PostForm("rf_key"), 10, 64)
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.RfKey = rfKey

	// 获取并转换值
	rfPower, err := postFormToInt(c, "rf_power")
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.RfPower = int8(rfPower)
	// 获取并转换值
	heartbeatCycle, err := postFormToInt(c, "heartbeat_cycle")
	if err != nil {
		// 发生错误，执行相应的处理
		log.Printf("Error occurred: %+v", err)
	}
	httpjson.HeartbeatCycle = int32(heartbeatCycle)
	// 从 map 中获取客户端对应的通道
	ch := DevClients[httpjson.DeviceMainIp]
	// 将字符串发送到通道

	json_p, err := json.Marshal(httpjson)
	if err != nil {
		log.Printf("转换json字符串失败！")
	}
	cmd := []byte{0xAA, 0x55, 0x00, 0xff, 0xF0, 0x00, 0x00, 0x00}

	ch <- append(cmd, json_p...)
}

func SetDeviceTimeSYNC(c *gin.Context) {
	addr := c.Query("device_ip")
	// 从 map 中获取客户端对应的通道
	ch := DevClients[addr]
	// 将字符串发送到通道

	for _, v := range DevCache.Items() {
		deviceReport := v.Object.(model.DeviceReportInfo)
		if strings.Split(deviceReport.DeviceIp, ":")[0] == addr {
			client := &http.Client{}
			// deviceReport.UTCtime[0] = 2023

			resp, err := client.Get(fmt.Sprintf("http://%s:8080/gopro/camera/set_date_time?date=%v_%v_%v&time=%v_%v_%v", deviceReport.GoProIP, deviceReport.UTCtime[0], deviceReport.UTCtime[1], deviceReport.UTCtime[2], deviceReport.UTCtime[3], deviceReport.UTCtime[4], deviceReport.UTCtime[5]))
			// fmt.Printf("http://%s:8080/gopro/camera/set_date_time?date=%v_%v_%v&time=%v_%v_%v", deviceReport.GoProIP, deviceReport.UTCtime[0], deviceReport.UTCtime[1], deviceReport.UTCtime[2], deviceReport.UTCtime[3], deviceReport.UTCtime[4], deviceReport.UTCtime[5])
			if err != nil {
				log.Printf("resp: %+v", resp)
				log.Printf("Error occurred: %+v", err)
				continue
			}
			defer resp.Body.Close()
			defer client.CloseIdleConnections()
		}
	}
	ch <- []byte{0xAA, 0x55, 0x00, 0xff, 0xF2, 0x00, 0x00, 0x00}
}

func DeviceControlWebcamStart(c *gin.Context) {
	// 发送第一个HTTP GET请求
	addr := c.Query("device_ip")
	goproIP := c.Query("gopro_ip")
	// 从 map 中获取客户端对应的通道
	ch := DevClients[addr]
	client := &http.Client{}
	resp, err := client.Get("http://" + goproIP + ":8080/gopro/camera/control/wired_usb?p=0")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)

		log.Printf("Error occurred: %+v", err)
	}

	// 将字符串发送到通道
	ch <- []byte{0xAA, 0x55, 0x00, 0xff, 0xF3, 0x00, 0x00, 0x00}

	c.JSONP(http.StatusOK, "ok")
}

func DeviceControlWebcamStop(c *gin.Context) {
	addr := c.Query("device_ip")
	goproIP := c.Query("gopro_ip")
	// 从 map 中获取客户端对应的通道
	ch := DevClients[addr]
	client := &http.Client{}
	resp, err := client.Get("http://" + goproIP + ":8080/gopro/camera/control/wired_usb?p=0")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}

	// 将字符串发送到通道
	ch <- []byte{0xAA, 0x55, 0x00, 0xff, 0xF4, 0x00, 0x00, 0x00}

	c.JSONP(http.StatusOK, "ok")
}

func GetDevList(c *gin.Context) {
	// c.JSONP(http.StatusOK, DevList)
}
