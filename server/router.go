package server

import (
	"gin-go/service"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.POST("/api/Login", service.Login)

	//全局中间件注册

	g.Use(service.Cors())

	s := g.Group("/system")
	{
		s.GET("/w", service.CmdWebSocket)
	}

	//Video Info
	v := g.Group("/api/video")
	{
		// d.Use(service.Auth)
		v.GET("/GetLocalVideoFileList", service.GetLocalVideoFileList)
		v.POST("/RenameLocalVideoFile", service.RenameLocalVideoFile)
		v.POST("/RemoveLocalVideoFile", service.RemoveLocalVideoFile)
		v.GET("/DownloadFile", service.DownloadFile)
		v.GET("/OpenFolder", service.OpenFolder)
		v.GET("/TimedVideoTask", service.TimedVideoTask)
		v.GET("/TimedVideoTaskCancel", service.TimedVideoTaskCancel)
	}

	//Device Info
	d := g.Group("/api/dev")
	{
		// d.Use(service.Auth)
		d.GET("/DelDev", service.DelDevice)
		d.GET("/GetDevList", service.GetDevList)
		d.GET("/SearchDeviceInfo", service.SearchDeviceInfo)
		d.POST("/SetDevice", service.SetDevice)
		d.GET("/DeviceControlWebcamStart", service.DeviceControlWebcamStart)
		d.GET("/DeviceControlWebcamStop", service.DeviceControlWebcamStop)
		d.GET("/SetDeviceTimeSYNC", service.SetDeviceTimeSYNC)

	}

	u := g.Group("/api/user")
	{
		// u.Use(service.Auth)
		// g.Use(service.CheckUserRole)

		u.POST("/RootCreateUser", service.RootCreateUser)
		u.GET("/GetUserList", service.GetUserList)
		u.POST("/RootDelUser", service.RootDelUser)
		u.POST("/RootResetPassWord", service.RootResetPassWord)
		u.POST("/RootSetUserAvailable", service.RootSetUserAvailable)
		u.POST("UserResetUserPassWord", service.UserResetUserPassWord)
		u.GET("/SetRoot", service.SetRoot)
	}

	log := g.Group("/api/log")
	{
		log.GET("/GetLogList", service.GetLogList)
		log.POST("/WriteLog", service.WriteLog)
		log.GET("/DelAllLog", service.DelAllLog)
	}

	Gopro := g.Group("/api/Gopro")
	{
		Gopro.GET("/GetGoProSate", service.GetGoProSate)
		Gopro.GET("/SetGoPro", service.SetGoPro)
		Gopro.GET("/GetGoProMediaList", service.GetGoProMediaList)
		Gopro.GET("/SetGoProShutter", service.SetGoProShutter)
		Gopro.GET("/SetGoProWebcam", service.SetGoProWebcam)
	}

	//地图文件服务
	g.StaticFS("/MapDownload", http.Dir("../MapDownload"))

	g.StaticFS("/VideoDownload", http.Dir("../VideoDownload"))

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//前端网页,把VUE2打包的dist 文件放到gin-go 目录下即可
	g.Use(static.Serve("/", static.LocalFile("dist", true)))
	g.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})

	return g
}
