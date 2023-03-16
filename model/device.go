package model

type DeviceReportInfo struct {
	DeviceName   string `json:"device_name"`   //设备名称，字符串
	DeviceIp     string `json:"device_ip"`     //设备IP
	GoProIP      string `json:"gopro_ip"`      //Gopro相机IP地址，字符串
	PosLatitude  int64  `json:"pos_latitude"`  //纬度，分扩大100000倍,实际要除以100000，整型
	PosNShemi    string `json:"pos_nshemi"`    //北纬/南纬，N:北纬，S:南纬，字符串
	PosLongitude int64  `json:"pos_longitude"` //经度，分扩大100000倍,实际要除以100000，整型
	PosWEhemi    string `json:"pos_ewhemi"`    //东经/西经，E:东经，W:西经，字符串
	UTCtime      []int  `json:"utc_time"`      //UTC时间，yyyy:年，mm:月，dd:日，hh:时，mm:分，ss:秒
	Battery      uint8  `json:"battery"`       //电池电量百分比，0~100，整型
	GoProState   uint8  `json:"gopro_state"`   //Gopro在线状态，0：不在线，1：在线
	RssiSnr      []int8 `json:"rssi_snr"`      //无线信号强度，整型，第一个是rssi，第二个是SNR
	//Extra        json.RawMessage `json:"extra"`
	// ReportNum int8 `json:"report_num"`
}

type DeviceSerchInfo struct {
	DeviceName     string `json:"device_name"`     //设备名称，字符串
	DeviceMainIp   string `json:"device_main_ip"`  //设备主IP，基于该IP与服务器通信，字符串
	DeviceSubIp    string `json:"device_sub_ip"`   //设备子IP，用于设备与GoPro通信，字符串
	DeviceGatWay   string `json:"device_gateway"`  //设备子IP，用于设备与GoPro通信，字符
	GoProIP        string `json:"gopro_ip"`        //Gopro相机IP地址，字符串
	VideoPort      int64  `json:"video_port"`      //设备给服务器转发视频端口，整型
	Baudrate       int64  `json:"baudrate"`        //设备串口波特率，整型
	RfModel        int8   `json:"rf_mode"`         //无线模式，0：中心节点，1：从节点
	RfBandwidth    int8   `json:"rf_bandwidth"`    //无线带宽，整型，有效值：3/5/10/20
	RfKey          int64  `json:"rf_key"`          //无线秘钥，整型有效值:100000~999999
	RfPower        int8   `json:"rf_power"`        //无线功率，整型，1：低，2：中，3：高
	HeartbeatCycle int32  `json:"heartbeat_cycle"` //心跳上报周期，整型，单位：秒
	TimeSyncCycle  int32  `json:"time_sync_cycle"` //北斗时间同步周期，整型，单位：秒
	// TimeSync       int8     `json:"time_sync"`       //时间同步控制，整型，1：设备执行时间同步，0：设备不执行时间同步，查询时该值始终为0
}
