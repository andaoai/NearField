package main

import (
	"gin-go/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	g := server.LoadRouter()
	server.ListenTCP()
	server.ListenControlTCP()
	// go func() {
	// 	log.Printf("启动360极速浏览器")
	// 	cmd := exec.Command("cmd", "/C", "start 360ChromeX --app= 127.0.0.1:8081")
	// 	cmd.Run() //自动打开
	// }()
	g.Run("0.0.0.0:8081")
}
