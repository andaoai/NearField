package service

import (
	"gin-go/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogList(c *gin.Context) {
	c.JSONP(http.StatusOK, dao.GetUserLogListDB())
}

func WriteLog(c *gin.Context) {
	act := c.PostForm("act")
	name := c.PostForm("name")
	role := c.PostForm("role")
	// userID, _ := strconv.Atoi(c.PostForm("userID"))
	dao.WriteLogDB(act, name, role)
}

func DelAllLog(c *gin.Context) {
	dao.DelAllLogDB()
}
