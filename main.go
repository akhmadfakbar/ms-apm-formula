package main

import (
	"github.com/akhmadfakbar/ms-apm-formula/connection"
	"github.com/akhmadfakbar/ms-apm-formula/handlers"
	"fmt"
)

func main() {
	connection.Connect()
	
	fmt.Println("Server up and running on port 11000")
    handlers.HandleRequests()
}