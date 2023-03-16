package model

type WebSocket struct {
	Data interface{} `json:"data"`
	Cmd  int         `json:"cmd"`
	//cmd:0 初始化列表信息
	//cmd:1 添加设备信息
	//cmd:2 设备下线，把列表中的设备在线状态标记为零
}
