package main

import (
	"fmt"
	"log"
)

func main() {

	db, err := NewPgClient()
	if err != nil {
		log.Fatal(err)
	}

	if  err := db.Initialize(); err != nil {
		log.Fatal(err)
	}  

	fmt.Println("Fired with account")

	server := NewEndPoint(":3000" , db)
	server.Fire()
}
