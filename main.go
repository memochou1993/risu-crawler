package main

import (
	"log"
	"net/http"

	"github.com/memochou1993/risu-crawler/controller"
)

func main() {
	go controller.Handle()

	log.Fatal(http.ListenAndServe(":7000", http.FileServer(http.Dir("storage"))))
}
