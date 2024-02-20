package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UserRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewTagsServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) repository.UsersService {
	return &UsersServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Buat implementasi TagsService
func (t *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	usersModel := model.Users{
		Username: users.Username,
	}
	t.UserRepository.Save(usersModel)
}

// Menghapus implementasi tagsService
func (t *UsersServiceImpl) Delete(usersID int) {
	t.UserRepository.Delete(usersID)
}

// Temukan semua implementasi
func (t *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UserRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			ID:       value.ID,
			Username: value.Username,
		}
		users = append(users, user)
	}

	return users
}

// Temukan berdasarkan id implementasi tagsService
func (t *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	usersData, err := t.UserRepository.FindById(usersId)
	helper.ErrorPanic(err)

	tagResponse := response.UsersResponse{
		ID:       usersData.UserDetails.UserID,
		Username: usersData.Username,
	}
	return tagResponse
}

// Mengupdate implementasi tagsService
func (t *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
	tagData, err := t.UserRepository.FindById(users.ID)
	helper.ErrorPanic(err)
	tagData.Username = users.Username
	t.UserRepository.Update(tagData)
}
