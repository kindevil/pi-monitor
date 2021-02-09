package websocket

import (
	"net/http"
	"pi-monitor/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type statistics struct {
	Host   *service.Host
	CPU    *service.CPU
	Memory *service.Memory
	Disks  []*service.Device
	NetDev []*service.Inter
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

var Conns []*websocket.Conn

func init() {
	go func() {
		for {
			if len(Conns) < 1 {
				time.Sleep(time.Second * 1)
				continue
			}

			st := &statistics{
				Host:   service.GetHost(),
				CPU:    service.GetCPU(),
				Memory: service.GetMem(),
				Disks:  service.GetDisk(),
				NetDev: service.GetNet(),
			}

			for _, conn := range Conns {
				if conn != nil {
					conn.WriteJSON(st)
				}
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	Conns = append(Conns, conn)

	//go Write(conn)

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		log.Infof("recv: %s", msgData)

		if msgType != websocket.TextMessage {
			continue
		}
	}
}

func Write(conn *websocket.Conn) {
	for {
		t1 := time.Now()
		st := &statistics{
			Host:   service.GetHost(),
			CPU:    service.GetCPU(),
			Memory: service.GetMem(),
			Disks:  service.GetDisk(),
			NetDev: service.GetNet(),
		}
		log.Info("timesince:", time.Since(t1))

		if conn != nil {
			conn.WriteJSON(st)
		}

		log.Info("sendto:", conn.RemoteAddr().String())

		time.Sleep(time.Millisecond * 100)
	}
}
