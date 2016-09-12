package game

import (
	"fmt"

	"github.com/k4s/tea/network"
	ms "github.com/k4s/teaServer/msg"
)

func MsgHello(msg interface{}, agent network.ExAgent) {
	m := msg.(*ms.Hello)
	fmt.Println("Hello,", m.Name)
	hi := &ms.Hello{Name: "kas", Age: 10}
	agent.WriteMsg(hi)
}

func MsgLogin(msg interface{}, agent network.ExAgent) {
	m := msg.(*ms.Login)
	if m.Name == "kas" && m.Password == "123456" {
		userdata := ms.UserData{}

		userdata.Account = m.Name
		userdata.Password = m.Password
		userdata.Auth = true
		agent.SetUserData(userdata)
		hi := &ms.Login{Name: "kas", Password: "123456"}
		agent.WriteMsg(hi)
		user := agent.UserData().(ms.UserData)
		fmt.Println("user:", user)
		return
	}
}
