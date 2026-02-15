package main

import (
	"fmt"
	router2 "go-fiber/router"
	"net/http"
)

func netflix() {
	fmt.Println("Hello world")
	router := router2.Router()
	http.ListenAndServe(":8000", router)
}
