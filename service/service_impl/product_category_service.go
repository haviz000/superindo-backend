package service_impl

import (
	"errors"
	"time"

	"github.com/haviz000/superindo-retail/helper"
	"github.com/haviz000/superindo-retail/model"
	"github.com/haviz000/superindo-retail/repository"
	"github.com/haviz000/superindo-retail/service"
)

type ProductCategoryServiceImpl struct {
	productCategoryRepository repository.ProductCategoryRepository
}

func NewProductCategoryService(categoryRepo repository.ProductCategoryRepository) service.ProductCategoryService {
	return &ProductCategoryServiceImpl{
		productCategoryRepository: categoryRepo,
	}
}

// GetOne implements service.ProductCategoryService.
func (pc *ProductCategoryServiceImpl) GetOne(categoryID string) (model.ProductCategoryResponse, error) {
	categoriesResult, err := pc.productCategoryRepository.FindByID(categoryID)
	if err != nil {
		return model.ProductCategoryResponse{}, err
	}

	return model.ProductCategoryResponse(categoriesResult), nil
}

// Create implements service.ProductCategoryService.
func (pc *ProductCategoryServiceImpl) Create(productCategoryReqdata model.ProductCategoryCreateRequest, userId string) (*model.ProductCategoryResponse, error) {
	categoryId := helper.GenerateID()
	newProductCategory := model.ProductCategory{
		ID:        categoryId,
		Name:      productCategoryReqdata.Name,
		Active:    "true",
		UserID:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := pc.productCategoryRepository.Create(newProductCategory)
	if err != nil {
		return nil, err
	}

	return &model.ProductCategoryResponse{
		ID:        newProductCategory.ID,
		Name:      newProductCategory.Name,
		Active:    newProductCategory.Active,
		UserID:    newProductCategory.UserID,
		CreatedAt: newProductCategory.CreatedAt,
		UpdatedAt: newProductCategory.UpdatedAt,
	}, nil
}

// Delete implements service.ProductCategoryService.
func (pc *ProductCategoryServiceImpl) Delete(productCategoryID string, userID string) (model.ProductCategoryResponse, error) {
	findProductCatgoryResponse, err := pc.productCategoryRepository.FindByID(productCategoryID)
	if err != nil {
		return model.ProductCategoryResponse{}, err
	}

	if userID != findProductCatgoryResponse.UserID {
		return model.ProductCategoryResponse{}, errors.New("Unauthorized")
	}

	err = pc.productCategoryRepository.Delete(model.ProductCategory{ID: productCategoryID})
	if err != nil {
		return model.ProductCategoryResponse{}, err
	}

	return model.ProductCategoryResponse{
		ID: productCategoryID,
	}, nil
}

// GetAll implements service.ProductCategoryService.
func (pc *ProductCategoryServiceImpl) GetAll() ([]model.ProductCategoryResponse, error) {
	productCategoryResult, err := pc.productCategoryRepository.FindAll()

	if err != nil {
		return []model.ProductCategoryResponse{}, err
	}

	productCategoryResponse := []model.ProductCategoryResponse{}

	for _, productCategoryRes := range productCategoryResult {
		productCategoryResponse = append(productCategoryResponse, model.ProductCategoryResponse(productCategoryRes))
	}

	return productCategoryResponse, nil
}

// Update implements service.ProductCategoryService.
func (pc *ProductCategoryServiceImpl) Update(productCategoryReqData model.ProductCategoryUpdateRequest, userID string, productCategoryId string) (*model.ProductCategoryResponse, error) {
	findProductCatgoryResponse, err := pc.productCategoryRepository.FindByID(productCategoryId)
	if err != nil {
		return nil, err
	}

	if userID != findProductCatgoryResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedCategoryReq := model.ProductCategory{
		ID:        productCategoryId,
		Name:      productCategoryReqData.Name,
		Active:    findProductCatgoryResponse.Active,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = pc.productCategoryRepository.Update(updatedCategoryReq)

	if err != nil {
		return nil, err
	}

	return &model.ProductCategoryResponse{
		ID: productCategoryId,
	}, nil
}
