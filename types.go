package main

import (
	"math/rand"	
	"time"
)

type BodyReq struct {
	FirstName string `json:"firstName"`
	LastName string`json:"lastName"`
}

type Account struct {
	firstName     string `json:"firstName"`
	lastName      string `json:"lastName"`
	accountNumber int    `json:"accountNumber"`
	balance       int64  `json:"balance"`
	created_at    time.Time `json:"created_at"`
}

// the abov struct with json annotation implies thatw when this data got serialized it will converted to respective labels in json

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		firstName:     firstName,
		lastName:      lastName,
		accountNumber: rand.Intn(10000),
		balance:       0,
		created_at:   time.Now().UTC(),

	}
}

type Logs struct {
	ID         int
	logDetails string
}
