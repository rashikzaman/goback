package domain

import (
	"context"
	"locally/goback/pkg/common/model"
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
	Fetch(ctx context.Context) []User
	FetchById(ctx context.Context)
	FetchByEmail(ctx context.Context)
	Store(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type UserUseCase interface {
	Fetch(ctx context.Context) []User
	FetchById(ctx context.Context)
	FetchByEmail(ctx context.Context)
	Store(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
