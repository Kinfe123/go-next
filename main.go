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

     fmt.Println("the db is pinging")
	fmt.Printf("%+v", db)

	server := NewEndPoint(":3000" , db)
	server.Fire()
}
