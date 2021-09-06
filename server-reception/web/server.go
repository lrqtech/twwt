package web

import (
	"fmt"
	"golang.org/x/net/websocket"
	"lrqtech/server/config"
	"net"
	"time"
)

func EchoServer(ws *websocket.Conn) {
	t, _ := time.ParseDuration(config.WebsocketDeadline)
	_ = ws.SetDeadline(time.Now().Add(t))
	conn, err := net.Dial("tcp", config.TcpAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go mWs(conn, ws)
	go mTcp(conn, ws)
	select {}
}
