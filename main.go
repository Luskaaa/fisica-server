package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade", err)
		return
	}

	defer conn.Close()
	log.Println("cliente conectado:", r.RemoteAddr)

	for {
		msgType, data, err := conn.ReadMessage()

		if err != nil {
			log.Println("desconectou:", err)
			break
		}
		log.Printf("recebi: %s", data)

		if err := conn.WriteMessage(msgType, data); err != nil {
			break
		}
	}

}

func main() {
	http.HandleFunc("/ws", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
