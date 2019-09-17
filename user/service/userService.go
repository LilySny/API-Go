package service

import (
	"../model"
	"../dto"
	"../dao"
	)


func FindByID(id int) *model.User {
	return dao.FindByID(id)
}

func FindByUsername(username string) *model.User {
	return dao.FindByUsername(username)
}

func FindAll() []*model.User {
	return dao.FindAll()
}

func Update(dto *dto.UserDto){
	dao.Update(dto)
}

func Save(createDto *dto.UserCreateDto) int {
	return dao.Save(createDto)
}

func Delete(id int) {
	dao.Delete(id)
}