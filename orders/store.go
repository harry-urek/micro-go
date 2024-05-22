package main

import "context"

type store struct {
	//add DB
}

func NewStore() *store {
	return &store{}

}

func (s *store) Create(context.Context) error {
	return nil

}
