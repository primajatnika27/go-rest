package mappers

import (
	"primajatnika/api-server/adapters/orm"
	"primajatnika/api-server/api/dto"
	"primajatnika/api-server/domain/entities"
)

func UserRequestModelToOrm(users *dto.UserRequestModel) *orm.UsersOrm {
	return &orm.UsersOrm{
		Name:     users.Name,
		Surname:  users.Surname,
		Username: users.Username,
		Password: users.Password,
		Email:    users.Email,
		Phone:    users.Phone,
		Company:  users.Company,
		Address:  users.Address,
		Active:   users.Active,
	}
}

func OrmToUserResponseModel(users *orm.UsersOrm) *entities.UserResponseModel {
	return &entities.UserResponseModel{
		ID:       users.ID,
		Name:     users.Name,
		Surname:  users.Surname,
		Username: users.Username,
		Password: users.Password,
		Email:    users.Email,
		Phone:    users.Phone,
		Company:  users.Company,
		Address:  users.Address,
		Active:   users.Active,
	}
}
