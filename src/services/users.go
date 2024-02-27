package services

import (
	"template/domain/entities"
	"template/domain/repositories"
)

type usersService struct {
	UsersRepository repositories.IUsersRepository
}

type IUsersService interface {
	GetAllUser() ([]entities.UserDataFormat, error)
	InsertNewAccount(data *entities.NewUserBody) bool
}

func NewUsersService(repo0 repositories.IUsersRepository) IUsersService {
	return &usersService{
		UsersRepository: repo0,
	}
}

func (sv usersService) GetAllUser() ([]entities.UserDataFormat, error) {
	userData, err := sv.UsersRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (sv usersService) InsertNewAccount(data *entities.NewUserBody) bool {
	status := sv.UsersRepository.InsertNewUser(data)
	return status
}
