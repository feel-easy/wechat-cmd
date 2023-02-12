package cmd

import (
	"fmt"

	"github.com/abiosoft/ishell/v2"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

type commander struct {
	bot   *openwechat.Bot
	self  *openwechat.Self
	shell *ishell.Shell
}

func (c *commander) Run() {
	c.shell.AddCmd(&ishell.Cmd{
		Name: "user",
		Help: "当前用户",
		Func: c.User,
	})
	c.shell.AddCmd(&ishell.Cmd{
		Name: "friends",
		Help: "好友",
		Func: c.Friends,
	})
	c.shell.AddCmd(&ishell.Cmd{
		Name: "groups",
		Help: "组",
		Func: c.Groups,
	})
	c.shell.Run()
}

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func NewCommander() *commander {
	shell := ishell.New()
	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode
	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 登陆
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
	}
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
	}
	// 去掉心跳的输出
	bot.SyncCheckCallback = func(resp openwechat.SyncCheckResponse) {}
	return &commander{
		bot:   bot,
		self:  self,
		shell: shell,
	}
}
