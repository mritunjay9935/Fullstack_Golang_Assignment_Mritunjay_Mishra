package api

import (
	"apica-assignment/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		return
	}
	defer ws.Close()

	h.cacheService.RegisterClient(ws)

	for {
		var cacheItem service.Item
		err := ws.ReadJSON(&cacheItem)
		if err != nil {
			log.Printf("error reading JSON from websocket: %v", err)
			break
		}
		h.cacheService.Set(cacheItem.Key, cacheItem.Value, time.Duration(cacheItem.Expiration)*time.Second)
	}
}

func (h *Handler) SendUpdatesToClients() {
	for {
		time.Sleep(2 * time.Second)
		items := h.cacheService.GetItems()
		for client := range h.cacheService.Clients {
			err := client.WriteJSON(items)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				h.cacheService.UnregisterClient(client)
			}
		}
	}
}
