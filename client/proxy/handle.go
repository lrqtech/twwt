package proxy

import (
	"fmt"
	"golang.org/x/net/websocket"
	"lrqtech/client/config"
	"net"
)

func mTcp(conn net.Conn, ws *websocket.Conn) {
	for {
		tBuf := make([]byte, config.Size)
		nT, err := conn.Read(tBuf)
		if err != nil {
			return
		}
		tData := tBuf[:nT]
		if len(tData) > 0 {
			writeTCPData(ws, tData)
		}
	}
}

func mWs(conn net.Conn, ws *websocket.Conn) {
	for {
		wBuf := make([]byte, config.Size)
		nWs, err := ws.Read(wBuf)
		if err != nil {
			return
		}
		wData := wBuf[:nWs]
		if len(wData) > 0 {
			writeWSData(conn, wData)
		}
	}
}

func writeWSData(conn net.Conn, data []byte) {
	_, _ = conn.Write(data)
	if config.Debug == true {
		fmt.Println(data)
	}
}

func writeTCPData(ws *websocket.Conn, data []byte) {
	_, _ = ws.Write(data)
	if config.Debug == true {
		fmt.Println(data)
	}
}
