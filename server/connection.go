package entrapped

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	ws   *websocket.Conn
	data chan []byte
}

func newConnection(rw http.ResponseWriter, req *http.Request) (*connection, error) {
	ws, wsErr := upgrader.Upgrade(rw, req, nil)
	if wsErr != nil {
		return nil, wsErr
	}

	return &connection{ws, make(chan []byte, 512)}, nil
}
