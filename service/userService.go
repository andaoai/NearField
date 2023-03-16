package service

import (
	"gin-go/dao"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetRoot(c *gin.Context) {
	dao.CreateRootDB()
}

func RootCreateUser(c *gin.Context) {
	rootName := c.PostForm("rootName")
	name := c.PostForm("name")
	password := c.PostForm("password")
	role := c.PostForm("role")
	available, err := strconv.ParseBool(c.PostForm("available"))
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}
	// fmt.Printf("name: %s; password: %s; role: %s", name, password, role)

	if dao.GetUserRoleDB(rootName) == "root" {
		dao.CreateUserDB(name, password, role, available)
	} else {
		c.JSON(200, gin.H{
			"message": "User permissions must be root",
		})
	}
}

func GetUserList(c *gin.Context) {
	c.JSONP(http.StatusOK, dao.GetUserListDB())
}

func RootDelUser(c *gin.Context) {

	id := c.PostForm("id")
	name := c.PostForm("name")
	intNum, _ := strconv.Atoi(id)

	if dao.GetUserRoleDB(name) == "root" {
		dao.DelUserDB(uint(intNum))
	} else {
		c.JSON(200, gin.H{
			"message": "User permissions must be root",
		})
	}
}

func RootSetUserAvailable(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	available := c.PostForm("available")
	intNum, _ := strconv.Atoi(id)
	boolAvailable, _ := strconv.ParseBool(available)
	if dao.GetUserRoleDB(name) == "root" {
		dao.SetUserAvailableDB(uint(intNum), boolAvailable)

	} else {
		c.JSON(200, gin.H{
			"message": "User permissions must be root",
		})
	}
}

func RootResetPassWord(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	password := c.PostForm("password")
	intNum, _ := strconv.Atoi(id)
	if dao.GetUserRoleDB(name) == "root" {
		dao.ResetPassWordDB(uint(intNum), password)
	} else {
		c.JSON(200, gin.H{
			"message": "User permissions must be root",
		})
	}
}

func UserResetUserPassWord(c *gin.Context) {

	name := c.PostForm("name")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")

	if dao.UserResetUserPassWordDB(name, oldPassword, newPassword) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "User set password successfully!",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Failed to set password for user!",
		})
	}

}
