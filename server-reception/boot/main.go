package boot

import (
	"fmt"
	"golang.org/x/net/websocket"
	"lrqtech/server/config"
	"lrqtech/server/web"
	"net/http"
	"os"
)

func Run() {
	fmt.Println("Starting websocket server ......")
	http.Handle(config.WebsocketPath, websocket.Handler(web.EchoServer))
	fmt.Println("Listen: ", config.ServerPort)
	fmt.Println("Websocket Path: ", config.WebsocketPath)
	fmt.Println("Websocket Deadline: ", config.WebsocketDeadline)
	fmt.Println("Minecraft Tcp Address: ", config.TcpAddress)
	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		println("ListenAndServe: ", err.Error())
		fmt.Println("Please try again")
		os.Exit(0)
	}
}
