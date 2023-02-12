package cmd

import "github.com/abiosoft/ishell/v2"

func (c *commander) Groups(ctx *ishell.Context) {
	// 获取所有的群组
	groups, err := c.self.Groups()
	if err != nil {
		ctx.Err(err)
	}
	ctx.Println(groups)
}
