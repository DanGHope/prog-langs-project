package main

import (
  "fmt"
  "net/http"
  "log"
)

//Connection port
const port = 5000

func main() {
  //start up hub for listening
  go myHub.run()

  http.HandleFunc("/ws",serveWs)
  http.Handle("/",http.FileServer(http.Dir("./public")))

	//Print to console
	fmt.Printf("Server Starting On Port: %d...\n", port)
	//Start Server on "port"
  err := http.ListenAndServe(":5000",nil)
  if err != nil {
    log.Fatal("ListenAndServe: ",err)
  }
}
