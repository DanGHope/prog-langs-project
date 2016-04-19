// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
"strings"
)

const (
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte

  //User data
  name string
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
  //if for loop ends close socket properly
	defer func() {
		myHub.unregister <- c
		closeSocket(c);
	}()
  c.ws.SetReadLimit(maxMessageSize)
	for {
		_, message, err := c.ws.ReadMessage()
    log.Println(string(message[:]))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
    msg := strings.Split(string(message[:]),";")
    log.Println(msg)
    switch(msg[0]){
      case "new-user":
        c.name = msg[1]
        myHub.broadcast <- message
      default:
        myHub.broadcast <- message
    }
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump() {
	defer func() {
		closeSocket(c)
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

func closeSocket(c* connection){
  msg := "disconnect;"
  msg += c.name
  myHub.broadcast <- []byte(msg)
  c.ws.Close();
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
  log.Println("CONNECTION")
  //Upgrade http request to Websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
  //Create new connection using previously created websocket
	c := &connection{send: make(chan []byte, 256), ws: ws}
	myHub.register <- c
	go c.writePump()
	c.readPump()
}
