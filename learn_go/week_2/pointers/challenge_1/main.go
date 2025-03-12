package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
	t_type := t.transactionType
	switch t_type {
	case transactionDeposit:
		(*c).balance += t.amount
		// return errors.New("deposit")
	case transactionWithdrawal:
		if (*c).balance < t.amount {
			return errors.New("insufficient funds")
		} else {
			(*c).balance -= t.amount
		}
		// return errors.New("withdrawal")
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}

// ?
