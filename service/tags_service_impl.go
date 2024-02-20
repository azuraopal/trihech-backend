package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewTagsServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) *UsersService {
	return &UsersService{
		userRepository: userRepository,
		Validate:       validate,
	}
}

// Buat implementasi TagsService
func (t *UserServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	usersModel := model.Users{
		Username: users.Username,
	}
	t.UsersRepository.Save(usersModel)
}

// Menghapus implementasi tagsService
func (t *UserServiceImpl) Delete(usersID int) {
	t.UsersRepository.Delete(usersID)
}

// Temukan semua implementasi
func (t *UserServiceImpl) FindAll() []response.UsersResponse {
	result := t.UsersRepository.FindAll()

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
func (t *UserServiceImpl) FindById(usersId int) response.UsersResponse {
	usersData, err := t.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	tagResponse := response.UsersResponse{
		ID:       usersData.UserDetails.UserID,
		Username: usersData.Username,
	}
	return tagResponse
}

// Mengupdate implementasi tagsService
func (t *UserServiceImpl) Update(users request.UpdateUsersRequest) {
	tagData, err := t.UsersRepository.FindById(users.ID)
	helper.ErrorPanic(err)
	tagData.Username = users.Username
	t.UsersRepository.Update(tagData)
}
