package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/superindo-retail/controller/controller_impl"
	"github.com/haviz000/superindo-retail/middleware"
	"github.com/haviz000/superindo-retail/repository/repository_impl"
	"github.com/haviz000/superindo-retail/service/service_impl"

	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {

	userRepository := repository_impl.NewUserRepository(db)
	categoryRepository := repository_impl.NewProductCategorylRepository(db)

	userService := service_impl.NewUserService(userRepository)
	categoryService := service_impl.NewProductCategoryService(categoryRepository)

	userController := controller_impl.NewUserController(userService)
	categoryController := controller_impl.NewProductCategoryController(categoryService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authUser := router.Group("/User", middleware.AuthMiddleware)
	{
		authUser.GET("/profile", userController.GetProfile)
		authUser.GET("/customers", userController.GetCustomers)
	}

	authCategory := router.Group("Category", middleware.AuthMiddleware)
	{
		authCategory.POST("", categoryController.Create)
		authCategory.GET("", categoryController.GetAll)
		authCategory.GET("/:category_id", categoryController.GetOne)
		authCategory.PUT("/:category_id", categoryController.Update)
		authCategory.DELETE("/:category_id", categoryController.Delete)
	}

}
