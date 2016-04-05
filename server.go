package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

//Connection port
const port = 5000

//User ...
type User struct {
	name string
}

var users []User

func disconnect(c echo.Context) error {
	name := c.Param("name")

	for i, element := range users {
		if element.name == name {
			users = append(users[:i], users[i+1:]...)
			fmt.Printf("%s Has Disconnected\n", name)
			return c.String(http.StatusOK, "Disconnected!!!")
		}
	}

	return c.String(http.StatusBadRequest, "You are not Connected!!!")
}

func connect(c echo.Context) error {
  name := c.Param("name")

	for i, element := range users {
		if element.name == name {
			return c.String(http.StatusBadRequest, "Cannot Connect "+name+" Is already connected!!!")
		}
    fmt.Printf("%d. %s\n", i, element)
	}

	users = append(users, User{name: name})
	fmt.Printf("%s Has Connected\n", name)
	return c.String(http.StatusOK, "Connected As: "+name)
}

func main() {
	//Create a new "echo" object
	e := echo.New()

	//Directory for static/public content
	e.Use(middleware.Static("public"))

	e.Get("/connect/:name", connect)
	e.Get("/disconnect/:name", disconnect)

	//Print to console
	fmt.Printf("Server Starting On Port: %d...\n", port)
	//Start Server on "port"
	e.Run(standard.New(":" + strconv.Itoa(port)))
}
