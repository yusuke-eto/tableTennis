package usecase

import (
	"tableTennis/domain"
	"tableTennis/domain/repository"
)

type UserUseCase interface {
	Insert(user domain.User) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) Insert(user domain.User) error {
	err := uu.userRepository.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
