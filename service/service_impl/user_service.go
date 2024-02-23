package service_impl

import (
	"errors"
	"fmt"

	"github.com/haviz000/superindo-retail/helper"
	"github.com/haviz000/superindo-retail/model"
	"github.com/haviz000/superindo-retail/repository"
	"github.com/haviz000/superindo-retail/service"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		userRepository: userRepo,
	}
}

func (us *UserServiceImpl) GetCustomers() ([]model.UserResponse, error) {
	customersResult, err := us.userRepository.FindAllCustomer()

	if err != nil {
		return []model.UserResponse{}, err
	}

	customersResponse := []model.UserResponse{}
	for _, customerResponse := range customersResult {
		customersResponse = append(customersResponse, model.UserResponse{
			UserID:    customerResponse.UserID,
			Username:  customerResponse.Username,
			Role:      customerResponse.Role,
			CreatedAt: customerResponse.CreatedAt,
			UpdatedAt: customerResponse.UpdatedAt})
	}

	return customersResponse, nil
}

func (us *UserServiceImpl) GetProfile(userID string) (model.User, error) {
	user, err := us.userRepository.FindByID(userID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		UserID:    userID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (us *UserServiceImpl) Login(userReqData model.UserLoginRequest) (*string, error) {
	userResponse, err := us.userRepository.FindByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.PasswordIsMatch(userReqData.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helper.GenerateAccessToken(userResponse)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (us *UserServiceImpl) Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	userID := helper.GenerateID()
	hashedPassword, err := helper.Hash(userReqData.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		UserID:   userID,
		Username: userReqData.Username,
		Password: hashedPassword,
		Role:     userReqData.Role,
	}

	err = us.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &model.UserRegisterResponse{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Password: newUser.Password,
		Role:     newUser.Role,
	}, nil
}
