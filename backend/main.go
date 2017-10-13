package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

var (
	port = 8080
)

func main() {
	router := vestigo.NewRouter()
	ch := make(chan byte, 1) //create channel with size of 1 byte, to prevent premature exit
	flag.IntVar(&port, "p", 8080, "int: port you'd like the go server to run on. Defaults to 8080")
	flag.Parse()

	// ROUTES
	// Serve Angular SPA from root ('/').
	// IMPORTANT: all Angular pages must be accessed from the app. I.e no directly navigating to victory page.
	router.Get("/*", http.FileServer(http.Dir("./dist")).ServeHTTP)
	// All API handlers are defined in handlers.go
	router.Get("/api/entries", listEntrants)
	router.Post("/api/entries", addEntrant)
	router.Put("/api/entries/:id", updateEntrant)
	router.Delete("/api/entries/:id", deleteEntrant)
	router.Get("/api/entries/:id", listEntrant)
	router.Get("/api/status/:id", entrantHasWon)

	portString := strconv.Itoa(port)
	// Start server in goroutine to prevent blocking
	go func() {
		log.Fatal("ListenAndServe:", http.ListenAndServe(":"+portString, router))
		ch <- 1
	}()

	fmt.Println("Listening on port :" + portString)
	<-ch // receive from channel, program won't exit until something gets received
}
