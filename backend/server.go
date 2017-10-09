package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

func main() {
	router := vestigo.NewRouter()
	ch := make(chan byte, 1) //create channel with size of 1 byte, to prevent premature exit

	// ROUTES
	router.Get("/*", http.FileServer(http.Dir("./dist")).ServeHTTP)

	// Start server in goroutine to prevent blocking
	go func() {
		log.Fatal("ListenAndServe:", http.ListenAndServe(":8080", router))
		ch <- 1
	}()

	fmt.Println("Listening on port :8080")
	<-ch // receive from channel, program won't exit until something gets received
}
