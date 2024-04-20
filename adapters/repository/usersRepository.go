package repository

import (
	"gorm.io/gorm"
	"primajatnika/api-server/adapters/orm"
)

type UsersRepository interface {
	Create(users *orm.UsersOrm) (*orm.UsersOrm, error)
}

type usersRepository struct {
	GormDB *gorm.DB
}

func (u usersRepository) Create(users *orm.UsersOrm) (*orm.UsersOrm, error) {
	result := u.GormDB.Create(&users)
	return users, result.Error
}

func NewUsersRepository(gormDB *gorm.DB) UsersRepository {
	return &usersRepository{
		GormDB: gormDB,
	}
}
