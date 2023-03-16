package service

import (
	"encoding/json"
	"gin-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// 通讯通道列表
var chanID int = 0

// 通讯通道列表
var ChanMap = make(map[int]chan model.WebSocket, 10)

func CmdWebSocket(c *gin.Context) {
	myChanID := chanID
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	Tcp2websocketch := make(chan model.WebSocket)
	ChanMap[chanID] = Tcp2websocketch
	chanID += 1

	defer func() {
		ws.Close()
		close(ChanMap[myChanID])
		delete(ChanMap, myChanID)
		log.Printf("ws Close!")
	}()

	//第一次执行
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Printf("Error occurred: %+v", err)

	}
	log.Printf("recv: %s", message)
	tmpdata := model.WebSocket{
		Data: DevCache.Items(),
		Cmd:  1,
	}
	bypeTmp, err := json.Marshal(&tmpdata)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}
	err = ws.WriteMessage(mt, bypeTmp)
	// fmt.Println(strtmp)
	if err != nil {
		log.Printf("Error occurred: %+v", err)
	}

	//持续接受新在线设备
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error occurred: %+v", err)
			break
		}

		log.Printf("recv: %s", message)

		// strtmp := <- Tcp2websocketch
		//收到TCP信息就通过通道发送到WEB前端去

		tmpdata, ok := <-Tcp2websocketch
		if !ok {
			log.Panicln(ok)
		}

		bypeTmp, err := json.Marshal(&tmpdata)
		if err != nil {
			log.Printf("Error occurred: %+v", err)
			break
		}
		err = ws.WriteMessage(mt, bypeTmp)
		// fmt.Println(strtmp)
		if err != nil {
			log.Printf("Error occurred: %+v", err)
			break
		}

	}
}
