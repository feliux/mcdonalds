package main

import "context"

type OrdersStore interface {
	Create(context.Context) error
}

type store struct {
	// our mongoDB here
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}
