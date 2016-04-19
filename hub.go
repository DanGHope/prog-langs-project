//Hub to deal with all the connections

package main

import (
"log"
)

//Struct for maintaining connection states
type hub struct {
	// Registered connections.
	connections map[*connection]bool
	// Inbound messages from the connections.
	broadcast chan []byte
	// Register requests from the connections.
	register chan *connection
	// Unregister requests from connections.
	unregister chan *connection
}

//initilize hub struct
var myHub = hub{
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

//Run infinitely looking for messages
func (myHub *hub) run() {
	for {
		select {

		//Register new connection
		case conn := <-myHub.register:
      conn.name = "Undefined"
      for otherUser := range myHub.connections{
        msg := "new-user;"
        msg+=otherUser.name
        conn.send <- []byte(msg)
      }
			myHub.connections[conn] = true

		//unregister connection
		case conn := <-myHub.unregister:
			if _, ok := myHub.connections[conn]; ok {
				delete(myHub.connections, conn)
				close(conn.send)
			}

		//Broadcast message to all connections
		case msg := <-myHub.broadcast:
			for conn := range myHub.connections {
				select {
				case conn.send <- msg:
        //close if bad message
				default:
          log.Println("Bad message")
					close(conn.send)
					delete(myHub.connections, conn)
				}
			}

		}
	}
}
