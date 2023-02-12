package main

import (
	"github.com/feel-easy/wechat-bot/cmd"
)

func main() {
	cmd := cmd.NewCommander()
	cmd.Run()
}
