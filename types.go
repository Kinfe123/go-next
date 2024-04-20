package main

import (
	"math/rand"
	"time"
)

type BodyReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type TransferReq struct {
	FromAccount int `json:"fromAccount"`
	ToAccount   int `json:"toAccount"`
	Amount      int `json:"amount"`
}
type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber int       `json:"accountNumber"`
	Balance       int64     `json:"balance"`
	Created_at    time.Time `json:"created_at"`
}

// the abov struct with json annotation implies thatw when this data got serialized it will converted to respective labels in json

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: rand.Intn(10000),
		Balance:       0,
		Created_at:    time.Now().UTC(),
	}
}

type Logs struct {
	ID         int
	logDetails string
}
