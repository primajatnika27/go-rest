package services

import (
	"primajatnika/api-server/adapters/mappers"
	"primajatnika/api-server/adapters/repository"
	"primajatnika/api-server/api/dto"
	"primajatnika/api-server/domain/entities"
)

type UserService interface {
	Create(users *dto.UserRequestModel) (*entities.UserResponseModel, error)
}

type userService struct {
	UserRepository repository.UsersRepository
}

func (u userService) Create(users *dto.UserRequestModel) (*entities.UserResponseModel, error) {
	maps := mappers.UserRequestModelToOrm(users)

	result, err := u.UserRepository.Create(maps)
	if err != nil {
		return nil, err
	}

	return mappers.OrmToUserResponseModel(result), nil
}

func NewUserService(userRepository repository.UsersRepository) UserService {
	return &userService{
		UserRepository: userRepository,
	}
}
