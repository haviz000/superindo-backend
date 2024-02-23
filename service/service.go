package service

import "github.com/haviz000/superindo-retail/model"

type UserService interface {
	Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error)
	Login(userReqData model.UserLoginRequest) (*string, error)
	GetProfile(userID string) (model.User, error)
	GetCustomers() ([]model.UserResponse, error)
}

type ProductCategoryService interface {
	Create(productCategoryReqdata model.ProductCategoryCreateRequest, userId string) (*model.ProductCategoryResponse, error)
	GetAll() ([]model.ProductCategoryResponse, error)
	GetOne(categoryID string) (model.ProductCategoryResponse, error)
	Update(productCategoryReqData model.ProductCategoryUpdateRequest, userId string, productCategoryId string) (*model.ProductCategoryResponse, error)
	Delete(productCategoryId string, userId string) (model.ProductCategoryResponse, error)
}
