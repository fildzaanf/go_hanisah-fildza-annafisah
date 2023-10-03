package usecase

import (
	"Praktikum/entity"
	"Praktikum/repository"
)

type UserUseCase interface {
	CreateUserController(user *entity.User) error
	GetUsersController() (err error, res interface{})
	LoginUserController(user *entity.User) (error, string)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(UserRepo repository.UserRepository) *userUseCase {

	return &userUseCase{
		userRepository: UserRepo,
	}
}

func (uc userUseCase) CreateUserController(user *entity.User) error {
	return uc.userRepository.CreateUserController(user)
}

func (uc userUseCase) GetUsersController() (err error, res interface{}) {
	return uc.userRepository.GetUsersController()
}
func (uc userUseCase) LoginUserController(user *entity.User) (error, string) {
	return uc.userRepository.LoginUserController(user)
}
