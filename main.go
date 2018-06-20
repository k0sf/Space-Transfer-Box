package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"runtime"
)

func main() {
	// os.Args = []string{"Space-Transfer-Box", "help"}
	app := cli.NewApp()
	// app.EnableBashCompletion = true
	app.HideHelp = false
	app.Name = "Space-Transfer-Box"
	// app.UsageText = "Space-Transfer-Box [options]"
	// app.ArgsUsage = "参数吗？"
	app.Usage = "空间转移盒 for " + runtime.GOOS + "/" + runtime.GOARCH
	app.Description = `Space-Transfer-Box  使用GO语言编写的proxy工具，为操作各种代理，提供实用功能。
	 具体功能, 参见 COMMANDS 列表
	 --------------------------------------------------------------------
	 https://github.com/YGoldking/Space-Transfer-Box 获取更多帮助信息!
	 https://github.com/YGoldking/Space-Transfer-Box/releases 获取程序更新信息!
	 --------------------------------------------------------------------
	 交流反馈:
		提交Issue: https://github.com/YGoldking/Space-Transfer-Box/issues`
	app.Version = "1.0.0"
	app.Author = "YGoldking/Space-Transfer-Box: https://github.com/YGoldking/Space-Transfer-Box"
	app.Authors = []cli.Author{
		{Name: "GoldKing", Email: "botnetxy@gmail.com"},
	}
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{Name: "name", Value: "bob", Usage: "a name to say"},
	// }
	app.Action = func(c *cli.Context) error {
		fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
		fmt.Println("Hello friend!")
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "shadowsocks2",
			Aliases: []string{"ss2"},
			Usage:   "Shadowsocks2 Module operation management",
			Description: `Shadowsocks - A secure socks5 proxy,designed to protect your Internet traffic.
			go-shadowsocks2 - A fresh implementation of Shadowsocks in Go.
	Features：
			 SOCKS5 proxy with UDP Associate
			 Support for Netfilter TCP redirect (IPv6 should work but not tested)
			 UDP tunneling (e.g. relay DNS packets)
			 TCP tunneling (e.g. benchmark with iperf3)`,
			// CustomHelpTemplate: "名称：hello\n 用法：撒旦法师打发",
			Action: func(c *cli.Context) error {
				fmt.Printf("？？？")
				return nil
			},
		},
		{
			Name:        "shadowsocksr",
			Aliases:     []string{"ssr"},
			Usage:       "ShadowsocksR Module operation management",
			Description: "ShadowsocksR - A fast tunnel proxy that helps you bypass firewalls.",
			Action: func(c *cli.Context) error {
				fmt.Printf("？？？")
				return nil
			},
		},
		{
			Name:        "outline",
			Aliases:     []string{"ol"},
			Usage:       "Outline Module operation management",
			Description: "Outline - use the popular Shadowsocks protocol.",
			Action: func(c *cli.Context) error {
				fmt.Printf("？？？")
				return nil
			},
		},
		{
			Name:        "v2ray",
			Aliases:     []string{"v2"},
			Usage:       "V2Ray Module operation management",
			Description: "V2Ray - The main role is to perform certain processing on the incoming network connection based on the user's configuration and send it to the specified server.",
			Action: func(c *cli.Context) error {
				fmt.Printf("？？？")
				return nil
			},
		},
		{
			Name: "brook",
			// Aliases:     []string{""},
			Usage:       "Brook Module operation management",
			Description: "Brook - Brook's goal is to reduce the configuration steps. Keep it simple, stupid.",
			Action: func(c *cli.Context) error {
				fmt.Printf("？？？")
				return nil
			},
		},
		{
			Name:        "switching",
			Aliases:     []string{"sc"},
			Usage:       "Switching Proxy (SS/SSR/V2Ray/Brook/Outline)",
			Description: "SS/SSR/V2Ray/Brook/Outline Switch between agents",
			Action: func(c *cli.Context) error {

				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:        "test",
					Usage:       "test hello",
					Description: "test xxxxxxxxxxxxxxxxx",
					Action: func(c *cli.Context) error {
						fmt.Println("test 子命令")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
