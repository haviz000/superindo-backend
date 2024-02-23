package repository_impl

import (
	"errors"
	"fmt"

	"github.com/haviz000/superindo-retail/model"
	"github.com/haviz000/superindo-retail/repository"
	"gorm.io/gorm"
)

type ProductCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductCategorylRepository(db *gorm.DB) repository.ProductCategoryRepository {
	return &ProductCategoryRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.ProductCategoryRepository.
func (pr *ProductCategoryRepositoryImpl) Create(productCategoryReqData model.ProductCategory) error {
	err := pr.DB.Create(&productCategoryReqData).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements repository.ProductCategoryRepository.
func (pr *ProductCategoryRepositoryImpl) Delete(productCategoryReqData model.ProductCategory) error {
	err := pr.DB.Delete(&productCategoryReqData).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements repository.ProductCategoryRepository.
func (pr *ProductCategoryRepositoryImpl) FindAll() ([]model.ProductCategory, error) {
	productCategories := []model.ProductCategory{}

	err := pr.DB.Find(&productCategories).Error
	if err != nil {
		return []model.ProductCategory{}, err
	}
	return productCategories, nil
}

// Update implements repository.ProductCategoryRepository.
func (pr *ProductCategoryRepositoryImpl) Update(productCategoryReqData model.ProductCategory) error {
	err := pr.DB.Save(&model.ProductCategory{
		ID:        productCategoryReqData.ID,
		Name:      productCategoryReqData.Name,
		UpdatedAt: productCategoryReqData.UpdatedAt,
		Active:    productCategoryReqData.Active,
		UserID:    productCategoryReqData.UserID,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

// FindByID implements repository.ProductCategoryRepository.
func (pr *ProductCategoryRepositoryImpl) FindByID(id string) (model.ProductCategory, error) {
	productCategory := model.ProductCategory{}

	fmt.Println("ini", id)

	err := pr.DB.Debug().Where("id = ?", id).Take(&productCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ProductCategory{}, err
		}

		return model.ProductCategory{}, err
	}

	return productCategory, nil
}
