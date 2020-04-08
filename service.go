package main

import (
	"errors"
)

// CustomerService provides operations on strings.
type CustomerService interface {
	CreateCustomer(name string) (string, error)
	// UpdateCustomer(string) int
	// ReadCustomer(string) int
	// DeleteCustomer(string) int
}

type customerService struct {
	// customers         customer.Repository
}

func (s *customerService) CreateCustomer(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	// s.customers.Store(name)
	return name, nil
}

// NewCustomerService creates a customer service with its dependencies.
func NewCustomerService( /** customers customer.Repository */ ) CustomerService {
	return &customerService{
		/** customers: customers, */
	}
}

// ErrEmpty error when empty params
var ErrEmpty = errors.New("empty name")
