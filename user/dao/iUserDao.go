package dao

import (
"../model"
"../dto"
)

type IUserDao interface {
	FindByID(id int) *model.User

	FindByUsername(username string) *model.User

	FindAll() []*model.User

	Update(dto *dto.UserDto)

	Save(createDto *dto.UserCreateDto) int

	Delete(id int)
}

