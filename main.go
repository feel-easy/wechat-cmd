package main

import (
	"github.com/feel-easy/wechat-cmd/cmd"
)

func main() {
	cmd := cmd.NewCommander()
	cmd.Run()
}
