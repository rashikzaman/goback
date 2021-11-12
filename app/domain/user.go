package domain

import (
	"context"
	"locally/goback/app/model"
)

type User struct {
	model.CommonModelFields
	Name         *string `json:"name" gorm:"type:varchar(255);not null"`
	Email        *string `json:"email" gorm:"unique;type:varchar(255);not null"`
	PhoneNumber  *string `json:"phone_number" gorm:"type:varchar(255);unique;not null"`
	Gender       string  `json:"gender"`
	Dob          string  `json:"dob"`
	Photo        string  `json:"photo"`
	IsVerified   *bool   `json:"is_verified" gorm:"default:false"`
	IsActive     *bool   `json:"is_active" gorm:"default:false"`
	IsSeller     *bool   `json:"is_seller" gorm:"default:false"`
	IsChatActive *bool   `json:"chat_active" gorm:"default:false"`
}

type UserRepository interface {
	FetchUsers(ctx context.Context) ([]User, error)
	FetchUserById(ctx context.Context)
	FetchUserByEmail(ctx context.Context)
	StoreUser(context.Context, *User) (*User, error)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
}

type UserUseCase interface {
	FetchUsers(ctx context.Context) (*model.Collection, error)
	FetchUserById(ctx context.Context)
	FetchUserByEmail(ctx context.Context)
	StoreUser(context.Context, *User) (*User, error)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
}
