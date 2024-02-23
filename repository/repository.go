package repository

import "github.com/haviz000/superindo-retail/model"

type UserRepository interface {
	Create(userRequestData model.User) error
	FindAllCustomer() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByID(id string) (model.User, error)
}

type ProductCategoryRepository interface {
	Create(productCategoryReqData model.ProductCategory) error
	FindAll() ([]model.ProductCategory, error)
	FindByID(id string) (model.ProductCategory, error)
	Update(productCategoryReqData model.ProductCategory) error
	Delete(productCategoryReqData model.ProductCategory) error
}
