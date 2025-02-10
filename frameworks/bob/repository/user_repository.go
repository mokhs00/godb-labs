package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"godb-labs/frameworks/bob/models"
	"godb-labs/frameworks/domain/entity"
	"godb-labs/frameworks/domain/repository"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	user, err := models.Users.Query(
		models.SelectWhere.Users.UserID.EQ(userID),
	).One(ctx, bob.New(r.db))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &entity.User{
		UserID:            user.UserID,
		Email:             user.Email,
		Name:              user.Name,
		PasswordEncrypted: user.PasswordEncrypted,
	}, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := models.Users.Query(
		models.SelectWhere.Users.Email.EQ(email),
	).One(ctx, bob.New(r.db))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &entity.User{
		UserID:            user.UserID,
		Email:             user.Email,
		Name:              user.Name,
		PasswordEncrypted: user.PasswordEncrypted,
	}, nil
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	savedUser, err := models.Users.Insert(&models.UserSetter{
		Email:             omit.From(user.Email),
		Name:              omit.From(user.Name),
		PasswordEncrypted: omit.From(user.PasswordEncrypted),
	}).One(ctx, bob.New(r.db))
	if err != nil {
		return nil, err
	}

	return &entity.User{
		UserID:            savedUser.UserID,
		Email:             savedUser.Email,
		Name:              savedUser.Name,
		PasswordEncrypted: savedUser.PasswordEncrypted,
	}, nil
}

func (r *UserRepository) Delete(ctx context.Context, user *entity.User) error {
	effected, err := models.Users.Delete(
		models.DeleteWhere.Users.UserID.EQ(user.UserID),
	).Exec(ctx, bob.New(r.db))
	if err != nil {
		return err
	}

	if effected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *UserRepository) DeleteByUserID(ctx context.Context, userID int64) error {
	effected, err := models.Users.Delete(
		models.DeleteWhere.Users.UserID.EQ(userID),
	).Exec(ctx, bob.New(r.db))
	if err != nil {
		return err
	}

	if effected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
