package user

import "github.com/jrmygp/user-management/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(id int) (models.User, error)
}
