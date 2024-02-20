package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(userrs request.UpdateUsersRequest)
	Delete(usersID int)
	FindById(usersID int) response.UsersResponse
	FindAll() []response.UsersResponse
}
