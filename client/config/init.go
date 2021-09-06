package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	TcpAddress      string
	WebsocketOrigin string
	WebsocketUrl    string
	Size            int
	Debug           bool
)

var cliTcpAddress = flag.String("a", "", "Input Your Tcp Address")
var cliWebsocketOrigin = flag.String("o", "", "Input Your Origin OF Websocket")
var cliWebsocketUrl = flag.String("u", "", "Input Your Websocket Url")
var cliSize = flag.Int("s", 216, "Byte Size \n Default: 216")
var cliDebug = flag.Bool("d", false, "Debug Mode \n Default: false")

func init() {
	flag.Parse()
	if *cliTcpAddress == "" || *cliWebsocketOrigin == "" || *cliWebsocketUrl == "" {
		fmt.Println("Param error")
		fmt.Println("Please try again")
		fmt.Println("Use 'client -h' for help")
		os.Exit(0)
	} else {
		TcpAddress = *cliTcpAddress
		WebsocketOrigin = *cliWebsocketOrigin
		WebsocketUrl = *cliWebsocketUrl
		Size = *cliSize
		Debug = *cliDebug
	}
}
