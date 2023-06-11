package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const ws_request_path = "/sakthimahendran/wireless/mouse/pad"

type WebServer struct {
	wsCon          *websocket.Conn
	mouseEventChan chan MouseEvent
}

func (ws *WebServer) Start(me chan MouseEvent) {
	ws.mouseEventChan = me

	http.HandleFunc(ws_request_path, ws.connectWebSocket)
	http.HandleFunc("/", ws.serveWebPage)

	http.ListenAndServe(":80", nil)
}

func (ws *WebServer) serveWebPage(w http.ResponseWriter, r *http.Request) {
	w.Write(webPage)
}

func (ws *WebServer) connectWebSocket(w http.ResponseWriter, r *http.Request) {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  256,
		WriteBufferSize: 0,
	}

	ws.wsCon, _ = wsUpgrader.Upgrade(w, r, nil)

	fmt.Println("Connected with ", r.RemoteAddr)

	go ws.startReadingWebSocket()
}

func (ws *WebServer) startReadingWebSocket() {
	var mouseEvent MouseEvent

	for {
		_, scroll, err := ws.wsCon.ReadMessage()

		if err != nil {
			continue
		}

		json.Unmarshal(scroll, &mouseEvent)

		ws.mouseEventChan <- mouseEvent
	}
}
