package main

import (
	"log"

	"github.com/MudassirDev/practice-server/server"
)

func main() {
	srv := server.CreateServer()

	log.Printf("server is listening at %v...\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
