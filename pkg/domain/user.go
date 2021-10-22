package domain

import (
	"context"
	"locally/goback/pkg/model"
	"time"
)

type User struct {
	model.CommonModelFields
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	Dob         time.Time `json:"dob"`
	Photo       string    `json:"photo"`
	IsVerified  bool      `json:"is_verified"`
	IsActive    bool      `json:"is_active"`
	IsSeller    bool      `json:"is_seller"`
	ChatActive  bool      `json:"chat_active"`
}

type UserRepository interface {
	FetchUsers(ctx context.Context) ([]User, error)
	FetchUserById(ctx context.Context)
	FetchUserByEmail(ctx context.Context)
	StoreUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
}

type UserUseCase interface {
	FetchUsers(ctx context.Context) (*model.Collection, error)
	FetchUserById(ctx context.Context)
	FetchUserByEmail(ctx context.Context)
	StoreUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
}
