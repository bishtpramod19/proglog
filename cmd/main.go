package main

import (
	"fmt"
	"log"

	"github.com/bishtpramod19/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	fmt.Println("Serving on port 8080")
	log.Fatal(srv.ListenAndServe())

}
