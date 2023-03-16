package service

import (
	"gin-go/dao"

	"github.com/gin-gonic/gin"
)

// var tokenList []string

// func generatorMD5(code string) string {
// 	MD5 := md5.New()
// 	_, _ = io.WriteString(MD5, code)
// 	return hex.EncodeToString(MD5.Sum(nil))
// }

func Login(c *gin.Context) {
	//SELECT * FROM users WHERE users.name = 'andao'
	name := c.PostForm("name")
	password := c.PostForm("password")
	if dao.LoginDB(name, password) {
		// 给客户端设置cookie
		// maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		// secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		// c.SetCookie("token", "123", 36000, "/", "localhost", false, true)
		// token := generatorMD5(name + time.Now().Format("2006-01-02 15:04:05"))
		// tokenList = append(tokenList, token)
		if dao.CheckUserAvailableDB(name) {
			c.JSON(200, gin.H{
				"message": "login is successful",
				"code":    "666",
				"role":    dao.GetUserRoleDB(name),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "user Available is false",
				"code":    "555",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "login is failed",
			"code":    "777",
		})
	}
}
