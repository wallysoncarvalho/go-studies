package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   float64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}

type CreatAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
