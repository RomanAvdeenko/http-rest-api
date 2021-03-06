package store

import "github.com/RomanAvdeenko/http-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(uint) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
