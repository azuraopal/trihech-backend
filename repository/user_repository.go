package repository

import "golang-crud-gin/model"

type UsersRepository interface {
	Save(user model.Users)
	Update(user model.Users)
	Delete(userID int)
	FindById(userID int) (tags model.Users, err error)
	FindAll() []model.Users
}
