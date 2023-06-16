package main

import (
	"alta.id/go-skeleton/product"
	"alta.id/go-skeleton/repo"
	"alta.id/go-skeleton/user"
)

func main() {
	s := &repo.Repo{}

	u := user.User{ID: 1, Name: "User 1"}
	p := product.Product{ID: 1, Name: "Product 1"}

	s.CreateUser(u)
	s.CreateProduct(p)

	s.PrintUsers()
	s.PrintProducts()
}
