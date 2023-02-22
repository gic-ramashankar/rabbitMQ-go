package main

import (
	"demo/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	log.Printf("Server started at port  : 9000 ")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}
