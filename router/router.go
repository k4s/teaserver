package router

import (
	"github.com/k4s/teaServer/game"
	"github.com/k4s/teaServer/msg"
	"github.com/k4s/teaServer/msg/process"
)

func init() {
	process.Processor.SetHandler(&msg.Login{}, game.MsgLogin)
}
