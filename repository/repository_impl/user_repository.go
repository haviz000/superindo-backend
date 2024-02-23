package repository_impl

import (
	"github.com/haviz000/superindo-retail/model"
	"github.com/haviz000/superindo-retail/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (ur *UserRepositoryImpl) Create(userReqData model.User) error {
	err := ur.DB.Create(&userReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepositoryImpl) FindAllCustomer() ([]model.User, error) {
	customers := []model.User{}

	err := ur.DB.Where("role = ?", "customer").Find(&customers).Error
	if err != nil {
		return []model.User{}, err
	}

	return customers, nil
}

func (ur *UserRepositoryImpl) FindByID(userID string) (model.User, error) {
	var user model.User

	err := ur.DB.First(&user, "user_id = ?", userID).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := ur.DB.First(&user, "username = ?", username).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
