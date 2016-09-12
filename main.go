package main

import (
	"github.com/k4s/tea"
	"github.com/k4s/teaserver/gate"
	_ "github.com/k4s/teaserver/router"
)

func main() {
	tea := tea.NewTea(gate.Gate)
	tea.Run()
}
