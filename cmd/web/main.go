package main

import (
	"log"

	"github.com/kokweikhong/maejiccode/internal/server"
	"github.com/kokweikhong/maejiccode/ui"
)

func main() {
	srv := server.New(ui.UiFS)
	
	log.Printf("Server starting on http://localhost:%s\n", srv.Port)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
