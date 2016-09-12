package conf

import (
	"time"
)

var (
	// gate conf
	Protocol          = "protobuf" //"json"
	MaxConnNum        = 100
	WritingNum        = 2000
	MaxMsgLen  uint32 = 16384 //8192

	//tcp
	TCPAddr      = "127.0.0.1:8080"
	LenMsgLen    = 2
	LittleEndian = false

	//websocket
	WSAddr      = "" //192.168.1.188:3563"
	HTTPTimeout = 10 * time.Second

	//log
	LogLevel = ""
	LogPath  = ""

	// other conf
	MysqlAddr = "root:@tcp(127.0.0.1:3306)/moneytree?charset=utf8&parseTime=true"
)
