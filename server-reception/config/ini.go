package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func loadIniFile() {
	cfg, err := ini.Load(IniFilePath)
	if err != nil {
		fmt.Println("Fail to read file: ", err)
		os.Exit(0)
	}
	TcpAddress = cfg.Section("minecraft").Key("TcpAddress").String()
	WebsocketPath = cfg.Section("server").Key("WebsocketPath").MustString("/ws")
	WebsocketDeadline = cfg.Section("server").Key("WebsocketDeadline").MustString("15s")
	ServerPort = cfg.Section("server").Key("ServerPort").String()
	Size = cfg.Section("server").Key("Size").MustInt(216)
	Debug = cfg.Section("mode").Key("Debug").MustBool(false)
	if TcpAddress == "" || ServerPort == "" {
		fmt.Println("Param 'TcpAddress' or 'ServerPort' in config file have some problem")
		fmt.Println("Please check and try again")
		os.Exit(0)
	}
}
