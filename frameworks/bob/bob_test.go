package bob

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"godb-labs/frameworks/bob/repository"
	"godb-labs/frameworks/domain/entity"
	"testing"
)

func TestBob(t *testing.T) {
	t.Run("UserRepository", func(t *testing.T) {
		ctx := context.Background()

		db, err := sql.Open("mysql", "dblabs:dblabs@tcp(localhost:3306)/dblabs?parseTime=true")
		if err != nil {
			t.Fatal(err)
		}

		userRepository := repository.NewUserRepository(db)

		saved, err := userRepository.Save(ctx, &entity.User{
			Email:             "Email",
			Name:              "Name",
			PasswordEncrypted: "PasswordEncrypted",
		})
		if err != nil {
			t.Fatal(err)
		}

		found, err := userRepository.FindUserByID(ctx, saved.UserID)
		if err != nil || found == nil {
			t.Fatal(err)
		}

		if err = userRepository.Delete(ctx, saved); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, saved, found)
	})
}
