package user

import "github.com/jrmygp/user-management/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	EditUser(user models.User) (models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateCheckIn(checkIn models.CheckIn) (models.CheckIn, error)
	UpdateCheckIn(checkIn models.CheckIn) (models.CheckIn, error)
	GetCheckInByID(id int) (models.CheckIn, error)
}
