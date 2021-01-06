package websocket

import (
	"net/http"
	"pi-monitor/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"
)

type statistics struct {
	Host   *service.Host
	CPU    *service.CPU
	Memory *service.Memory
	Net    *service.Net
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

var Conn *websocket.Conn

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	defer conn.Close()

	logger.Debug(conn)

	Conn = conn

	go Write()

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			switch err.(type) {
			case *websocket.CloseError:
				Conn = nil
				return
			default:
				Conn = nil
				return
			}
		}

		if msgType != websocket.TextMessage {
			continue
		}

		logger.Info("incoming message: %s\n", msgData)
	}
}

func Write() {
	for {
		st := &statistics{
			Host:   service.GetHost(),
			CPU:    service.GetCPU(),
			Memory: service.GetMem(),
			Net:    service.GetNet(),
		}

		if Conn != nil {
			Conn.WriteJSON(st)
		}

		time.Sleep(time.Second * 1)
	}
}
