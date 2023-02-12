package cmd

import "github.com/abiosoft/ishell/v2"

func (c *commander) Friends(ctx *ishell.Context) {
	// 获取所有的好友
	friends, err := c.self.Friends()
	ctx.Println(friends, err)
}
