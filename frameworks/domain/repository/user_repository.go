package repository

import (
	"context"
	"godb-labs/frameworks/domain/entity"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, userID int64) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, user *entity.User) error
	DeleteByUserID(ctx context.Context, userID int64) error
}
