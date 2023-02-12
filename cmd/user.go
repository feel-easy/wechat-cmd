package cmd

import "github.com/abiosoft/ishell/v2"

func (c *commander) User(ctx *ishell.Context) {
	self, err := c.bot.GetCurrentUser()
	if err != nil {
		ctx.Err(err)
		return
	}
	c.self = self
	ctx.Printf("%v\n", self)
}
