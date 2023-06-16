package repo

import (
	"fmt"

	"alta.id/go-skeleton/product"
	"alta.id/go-skeleton/user"
)

type Repo struct {
	users    []user.User
	products []product.Product
}

func (s *Repo) CreateUser(u user.User) {
	s.users = append(s.users, u)
}

func (s *Repo) CreateProduct(p product.Product) {
	s.products = append(s.products, p)
}

func (s *Repo) PrintUsers() {
	for _, u := range s.users {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}
}

func (s *Repo) PrintProducts() {
	for _, p := range s.products {
		fmt.Printf("ID: %d, Name: %s\n", p.ID, p.Name)
	}
}
