package gate

import (
	GT "github.com/k4s/tea/gate"
	"github.com/k4s/teaServer/msg/process"
	"github.com/k4s/teaserver/conf"
)

var Gate = &GT.Gate{
	MaxConnNum:   conf.MaxConnNum,
	WritingNum:   conf.WritingNum,
	MaxMsgLen:    conf.MaxMsgLen,
	WSAddr:       conf.WSAddr,
	HTTPTimeout:  conf.HTTPTimeout,
	TCPAddr:      conf.TCPAddr,
	LenMsgLen:    conf.LenMsgLen,
	LittleEndian: conf.LittleEndian,
	Processor:    process.Processor,
}
