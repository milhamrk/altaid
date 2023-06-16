package store

import (
	"fmt"

	"alta.id/go-skeleton/product"
	"alta.id/go-skeleton/user"
)

type Store struct {
	users    []user.User
	products []product.Product
}

func (s *Store) CreateUser(u user.User) {
	s.users = append(s.users, u)
}

func (s *Store) CreateProduct(p product.Product) {
	s.products = append(s.products, p)
}

func (s *Store) PrintUsers() {
	for _, u := range s.users {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}
}

func (s *Store) PrintProducts() {
	for _, p := range s.products {
		fmt.Printf("ID: %d, Name: %s\n", p.ID, p.Name)
	}
}
