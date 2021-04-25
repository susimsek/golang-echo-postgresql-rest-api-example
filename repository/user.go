package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"golang-echo-postgresql-rest-api-example/exception"
	"golang-echo-postgresql-rest-api-example/model"
	"golang-echo-postgresql-rest-api-example/util"
)

var cntx context.Context = context.TODO()

type UserRepository interface {
	Count() int64
	GetAllUser(page int64, limit int64) (*util.PagedModel, error)
	SaveUser(user *model.User) (*model.User, error)
	GetUser(id string) (*model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (userRepository *userRepositoryImpl) Count() int64 {
	var count int64
	userRepository.db.Model(&model.User{}).Count(&count)
	return count
}

func (userRepository *userRepositoryImpl) GetAllUser(page int64, limit int64) (*util.PagedModel, error) {
	cc := make(chan int64, 0)
	var users []model.User

	go countRecords(userRepository.db, &users, cc)
	count := <-cc
	paginator := util.Paging(page, limit, count)

	err := userRepository.db.Model(&model.User{}).Limit(paginator.Limit).Offset(paginator.Offset).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return paginator.PagedData(users), nil
}

func (userRepository *userRepositoryImpl) SaveUser(user *model.User) (*model.User, error) {
	err := userRepository.db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepositoryImpl) GetUser(id string) (*model.User, error) {
	var existingUser model.User

	err := userRepository.db.Model(&model.User{}).First(&existingUser, "id = ?", id).Error
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	var existingUser model.User

	err := userRepository.db.Model(&model.User{}).First(&existingUser, "id = ?", id).Error
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	existingUser.UserInput = user.UserInput

	err = userRepository.db.Model(&model.User{}).Save(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) DeleteUser(id string) error {

	result := userRepository.db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return exception.ResourceNotFoundException("User", "id", id)
	}

	return nil
}

func countRecords(db *gorm.DB, anyType interface{}, cc chan int64) {
	var count int64
	db.Model(anyType).Count(&count)
	cc <- count
}
