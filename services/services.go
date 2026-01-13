package services

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var rooms = make(map[string]map[*websocket.Conn]bool)

var H1 = func(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Hello World")
}

var HandleWebSocket = func(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgader.Upgrade(w, r, nil)
	vars := mux.Vars(r)
	fileID := vars["fileId"]

	log.Printf("User connecting to file: %s", fileID)

	if rooms[fileID] == nil { // we check if the room doesn't already exist by checking if its nil(null)
		// Create the inner map
		rooms[fileID] = make(map[*websocket.Conn]bool)
	}
	rooms[fileID][connection] = true // we add the user to the room

}
