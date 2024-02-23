package controller_impl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/superindo-retail/controller"
	"github.com/haviz000/superindo-retail/helper"
	"github.com/haviz000/superindo-retail/model"
	"github.com/haviz000/superindo-retail/service"
)

type ProductCategoryControllerImpl struct {
	ProductCategoryService service.ProductCategoryService
}

func NewProductCategoryController(service service.ProductCategoryService) controller.ProductCategoryController {
	return &ProductCategoryControllerImpl{
		ProductCategoryService: service,
	}
}

// Create implements controller.ProductCategoryController.
func (pc *ProductCategoryControllerImpl) Create(ctx *gin.Context) {
	var request model.ProductCategoryCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.ProductCategoryCreateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := pc.ProductCategoryService.Create(request, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "product category created successfully",
		Data:    response,
	})
}

// Delete implements controller.ProductCategoryController.
func (pc *ProductCategoryControllerImpl) Delete(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := pc.ProductCategoryService.Delete(categoryID, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "category deleted successfully",
		Data: model.ProductCategoryResponse{
			ID: response.ID,
		},
	})
}

// GetAll implements controller.ProductCategoryController.
func (pc *ProductCategoryControllerImpl) GetAll(ctx *gin.Context) {
	response, err := pc.ProductCategoryService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get all category successfully",
		Data:    response,
	})
}

// GetOne implements controller.ProductCategoryController.
func (pc *ProductCategoryControllerImpl) GetOne(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")

	response, err := pc.ProductCategoryService.GetOne(categoryID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get category successfully",
		Data:    response,
	})
}

// Update implements controller.ProductCategoryController.
func (pc *ProductCategoryControllerImpl) Update(ctx *gin.Context) {
	var request model.ProductCategoryUpdateRequest
	categoryID := ctx.Param("category_id")
	fmt.Println("ksong", categoryID)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.CategoryUpdateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := pc.ProductCategoryService.Update(request, userID.(string), categoryID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "category updated successfully",
		Data: model.ProductCategoryUpdateResponse{
			ID: response.ID,
		},
	})
}
