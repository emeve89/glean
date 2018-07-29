package usecases

import "github.com/emeve89/glean/domain"

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}

type User struct {
	Id 			int
	IsAdmin 	bool
	Customer 	domain.Customer
}
