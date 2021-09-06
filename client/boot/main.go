package boot

import (
	"fmt"
	"golang.org/x/net/websocket"
	"lrqtech/client/config"
	"lrqtech/client/proxy"
	"net"
	"net/http"
	"net/url"
	"os"
)

func Run() {
	fmt.Println("Starting tcp server ......")
	listener, err := net.Listen("tcp", config.TcpAddress)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return
	}
	fmt.Println("Tcp listen at: ", config.TcpAddress)
	fmt.Println("Dial: ", config.WebsocketUrl)
	c := new(websocket.Config)
	c.Version = websocket.ProtocolVersionHybi13
	c.Location, err = url.ParseRequestURI(config.WebsocketUrl)
	if err != nil {
		return
	}
	c.Origin, err = url.ParseRequestURI(config.WebsocketOrigin)
	if err != nil {
		return
	}
	h := http.Header(make(map[string][]string))
	h.Add("User-Agent", "Go-http-client/1.1")
	c.Header = h
	ws, err := websocket.DialConfig(c)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Please try again")
		os.Exit(0)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go proxy.Proxy(conn, ws)
	}
}
