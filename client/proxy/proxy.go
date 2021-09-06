package proxy

import (
	"golang.org/x/net/websocket"
	"net"
)

func Proxy(conn net.Conn, ws *websocket.Conn) {
	go mTcp(conn, ws)
	go mWs(conn, ws)
	select {}
}
