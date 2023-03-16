package dao

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"UNIQUE"`
	PassWord  string
	Role      string
	Available bool
	// OnlineTime time.Time
}
type UserLog struct {
	gorm.Model
	Act  string
	Name string
	Role string
}

func dbinit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sys.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Error occurred: %+v", err)
		// log.Printf("Error occurred: %+v", err)
		panic("failed to connect database")
	}
	//
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserLog{})
	return db
}

var db *gorm.DB = dbinit()

func CreateUserDB(name string, pw string, role string, available bool) {
	// id := common.GeneratorMD5(name + time.Now().Format("2006-01-02 15:04:05") + name)
	db.Create(&User{
		Name:      name,
		PassWord:  pw,
		Role:      role,
		Available: available,
	})
}

func CreateRootDB() {
	// id := common.GeneratorMD5(name + time.Now().Format("2006-01-02 15:04:05") + name)
	db.Create(&User{
		Name:      "root",
		PassWord:  "123456",
		Role:      "root",
		Available: true,
	})
}

func GetUserListDB() (Userlist []User) {
	db.Raw("SELECT users.id,users.created_at,users.name,users.role,users.available FROM users WHERE users.deleted_at IS NULL").Scan(&Userlist)
	return
}

func ResetPassWordDB(id uint, newPW string) {
	var user User
	db.First(&user, id)
	db.Model(&user).Update("pass_word", newPW)
}

func UserResetUserPassWordDB(name string, oldPW string, newPW string) bool {
	var user User
	db.First(&user, "name = ?", name)
	if user.PassWord == oldPW {
		db.Model(&user).Update("pass_word", newPW)
		return true
	} else {
		return false
	}

}

func DelUserDB(id uint) {
	db.Delete(&User{}, id)
}

func WriteLogDB(act string, name string, role string) {
	db.Create((&UserLog{
		Act:  act,
		Name: name,
		Role: role,
	}))
}

func DelAllLogDB() {
	db.Where("1 = 1").Unscoped().Delete(&UserLog{})
}

func GetUserLogListDB() (Loglist []UserLog) {
	db.Raw("SELECT * FROM user_logs WHERE user_logs.deleted_at IS NULL order by id desc").Scan(&Loglist)
	return
}

func CheckUserAvailableDB(name string) bool {
	var user User
	//SELECT users.available FROM users WHERE users.deleted_at IS NULL and users.name = 'andao'
	// db.Raw("SELECT * FROM user_logs WHERE user_logs.deleted_at IS NULL").Scan(&Loglist)
	db.First(&user, "name = ?", name)
	return user.Available
}

func GetUserRoleDB(name string) string {
	var user User
	db.First(&user, "name = ?", name)
	return user.Role
}

func SetUserAvailableDB(id uint, available bool) {
	var user User
	db.First(&user, id)
	db.Model(&user).Update("available", available)
}

func LoginDB(name string, pw string) bool {
	var user User

	if name == "" || pw == "" {
		return false
	}

	db.First(&user, "name = ?", name)
	return pw == user.PassWord
}

//https://blog.csdn.net/cnwyt/article/details/118904882?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-0-118904882-blog-122519180.pc_relevant_multi_platform_whitelistv3&spm=1001.2101.3001.4242.1&utm_relevant_index=2
// func Test() {
// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})

// 	// 插入内容
// 	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
// 	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

// 	// 读取内容
// 	var product Product
// 	db.First(&product, 1)                 // find product with integer primary key
// 	db.First(&product, "code = ?", "D42") // find product with code D42

// 	// 更新操作： 更新单个字段
// 	db.Model(&product).Update("Price", 2000)

// 	// 更新操作： 更新多个字段
// 	db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
// 	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})

// 	// 删除操作：
// 	db.Delete(&product, 1)
// }
