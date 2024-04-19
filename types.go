package main

import (
	"math/rand"
)

type BodyReq struct {
	FirstName string `json:"firstName"`
	LastName string`json:"lastName"`
}

type Account struct {
	ID            int    `json:"id"`
	firstName     string `json:"firstName"`
	lastName      string `json:"lastName"`
	accountNumber int    `json:"accountNumber"`
	balance       int64  `json:"balance"`
}

// the abov struct with json annotation implies thatw when this data got serialized it will converted to respective labels in json

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:            rand.Intn(1000),
		firstName:     firstName,
		lastName:      lastName,
		accountNumber: rand.Intn(10000),
		balance:       0,
	}
}

type Logs struct {
	ID         int
	logDetails string
}
