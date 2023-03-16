package service

import (
	"gin-go/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// 获取客户端cookie并校验
	if cookie, err := c.Cookie("token"); err == nil {
		if cookie == "123" {
			c.Next()
			return
		}
	}
	// 返回错误
	c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
	// 若验证不通过，不再调用后续的函数处理
	c.Abort()
}

func CheckUserRole(c *gin.Context) {

	name := c.Query("name")
	if !dao.CheckUserAvailableDB(name) {
		// log.Printf("test")
		//abort（）顾名思义就是终止的意思，也就是说执行该函数，会终止后面所有的该请求下的函数。
		c.Abort()
	}
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
