package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoProSate(c *gin.Context) {
	ip := c.Query("ip")
	client := &http.Client{}

	// 发送第一个HTTP GET请求，开启相机的USB控制
	resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=1")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}

	// 发送第二个HTTP GET请求，查询相机状态
	resp, err = client.Get("http://" + ip + ":8080/gopro/camera/state")
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	// 使用defer语句关闭resp.Body
	defer resp.Body.Close()

	// 读取resp.Body中的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	// 将JSON数据解析为map
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(string(body)), &data)

	// 返回JSONP格式的响应
	c.JSONP(http.StatusOK, data)
}

func SetGoPro(c *gin.Context) {
	ip := c.Query("ip")
	setting := c.Query("setting")
	option := c.Query("option")
	client := &http.Client{}
	resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=1")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}
	resp, err = client.Get("http://" + ip + ":8080/gopro/camera/setting?setting=" + setting + "&option=" + option)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}
	var tmp map[string]interface{}
	_ = json.Unmarshal([]byte(string(body)), &tmp)
	c.JSONP(http.StatusOK, tmp)
	// log.Printf(string(body))
}

func GetGoProMediaList(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSONP(http.StatusBadRequest, gin.H{"error": "ip is required"})
		return
	}
	client := &http.Client{}
	resp, err := client.Get("http://" + ip + ":8080/gopro/media/list")
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, result)
}

func SetGoProShutter(c *gin.Context) {
	ip := c.Query("ip")
	q := c.Query("q")
	client := &http.Client{}

	// 发送第一个HTTP GET请求
	resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=1")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		// log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}

	// 发送第二个HTTP GET请求
	resp, err = client.Get("http://" + ip + ":8080/gopro/camera/presets/set_group?id=1000")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}

	// 发送第三个HTTP GET请求
	resp, err = client.Get("http://" + ip + ":8080/gopro/camera/shutter/" + q)
	log.Printf("http://" + ip + ":8080/gopro/camera/shutter/" + q)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	resp, err = client.Get("http://" + ip + ":8080/gopro/camera/state")
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	// 使用defer语句关闭resp.Body
	defer resp.Body.Close()

	// 读取resp.Body中的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	// 将JSON数据解析为map
	var data map[string]map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &data)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}
	log.Printf("test: %+v", data["status"]["8"])
	log.Printf("test: %+v", data["status"]["10"])
	// 使用defer语句关闭resp.Body
	defer resp.Body.Close()
}

func SetGoProWebcam(c *gin.Context) {
	ip := c.Query("ip")
	q := c.Query("q")
	client := &http.Client{}

	// 发送第一个HTTP GET请求
	resp, err := client.Get("http://" + ip + ":8080/gopro/camera/control/wired_usb?p=0")
	if err != nil {
		log.Printf("Error occurred: %+v", resp)
		log.Printf("Error occurred: %+v", err)
	}

	// 发送第二个HTTP GET请求
	resp, err = client.Get("http://" + ip + ":8080/gopro/webcam/" + q)
	log.Printf("http://" + ip + ":8080/gopro/webcam/" + q)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	log.Printf(ip + "," + q)
	// 使用defer语句关闭resp.Body
	defer resp.Body.Close()
}
