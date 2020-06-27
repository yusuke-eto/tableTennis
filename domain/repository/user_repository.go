package repository

import (
	"tableTennis/domain"
)

// UserRepository is interface
type UserRepository interface {
	Insert(user domain.User) error
}
