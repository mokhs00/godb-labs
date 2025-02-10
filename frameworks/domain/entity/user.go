package entity

type User struct {
	UserID            int64
	Email             string
	Name              string
	PasswordEncrypted string
}
