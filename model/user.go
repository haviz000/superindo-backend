package model

import "time"

type User struct {
	UserID    string `gorm:"primaryKey;type:varchar(255)"`
	Username  string `gorm:"not null;uniqueIndex;type:varchar(50)"`
	Password  string `gorm:"not null;type:varchar(255)"`
	Role      string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserRegisterResponse struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
