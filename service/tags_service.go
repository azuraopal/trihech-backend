package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type UsersService interface {
	Create(tags request.CreateUsersRequest)
	Update(tags request.UpdateUsersRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
