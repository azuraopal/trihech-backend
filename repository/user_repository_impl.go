package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Menghapus implementasi TagsRepository
func (t *UsersRepositoryImpl) Delete(userID int) {
	var user model.Users
	result := t.Db.Where("id = ?", userID).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// temukan semua implementsasi TagsRepository
func (t *UsersRepositoryImpl) FindAll() []model.Users {
	var userFindAll []model.Users
	result := t.Db.Find(&userFindAll)
	helper.ErrorPanic(result.Error)
	return userFindAll
}

// menemukan melalui id implementasi UsersRepository
func (t *UsersRepositoryImpl) FindById(usersID int) (user model.Users, err error) {
	var users model.Users
	result := t.Db.Find(&users, usersID)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user is not found")
	}
}

// Menyimpan implementasi UsersRepository
func (t *UsersRepositoryImpl) Save(user model.Users) {
	result := t.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Mengupdate implementasi TagsRepository
func (t *UsersRepositoryImpl) Update(users model.Users) {
	var updateUser = request.UpdateUsersRequest{
		ID:   int(users.ID),
		Username: users.Username,
	}
	result := t.Db.Model(&users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
