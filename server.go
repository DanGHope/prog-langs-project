package main

/*
import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"strconv"
	"strings"
)

//User ...
type User struct {
	ws   *websocket.Conn
	name string
}

//Message for websocket
type Message struct {
	mType string
	name  string
}

var users []User

func serveUser(ws *websocket.Conn) {
	for {
		msg := ""
		websocket.Message.Receive(ws, &msg)
		data := strings.Split(msg, ";")

		if data[0] == "connect" {
			users = append(users, User{name: data[1], ws: ws})
			for _, element := range users {
				websocket.Message.Send(ws, "new-user;"+element.name)
				if element.ws != ws {
					websocket.Message.Send(element.ws, "new-user;"+data[1])
				}
			}
		}

		if data[0] == "disconnect" {
			for _, element := range users {
				websocket.Message.Send(element.ws, "disconnect;"+data[1])
			}
			return
		}
	}
}

func disconnect(c echo.Context) error {
	name := c.Param("name")

	for i, element := range users {
		if element.name == name {
			users = append(users[:i], users[i+1:]...)
			fmt.Printf("%s Has Disconnected\n", name)
			return c.String(http.StatusOK, "Disconnected!")
		}
	}

	return c.String(http.StatusOK, "You are not Connected!!!")
}

func connect(c echo.Context) error {
	name := c.Param("name")

	for _, element := range users {
		if element.name == name {
			return c.String(http.StatusOK, "Cannot Connect "+name+" Is already connected!!!")
		}
	}

	fmt.Printf("%s Has Connected\n", name)
	return c.String(http.StatusOK, "Connected As: "+name)
}
*/
