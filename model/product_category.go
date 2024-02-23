package model

import "time"

type ProductCategory struct {
	ID        string `gorm:"primaryKey;type:varchar(255)"`
	Name      string `gorm:"not null;type:varchar(50)"`
	Active    string `gorm:"not null;type:boolean"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductCategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Active    string    `json:"active"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductCategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductCategoryUpdateRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductCategoryUpdateResponse struct {
	ID string `json:"id"`
}

type ProductCategoryDeleteResponse struct {
	ID string `json:"id"`
}
