// websocket/client.go

package websocket

import (
	gws "github.com/gorilla/websocket"
)

type Client struct {
	Conn *gws.Conn
	Send chan any
}
