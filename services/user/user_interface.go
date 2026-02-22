package user

import (
	"github.com/jrmygp/user-management/models"
	"github.com/jrmygp/user-management/requests"
)

type UserService interface {
	CreateUser(user requests.CreateUserRequest) (models.User, error)
	GetUserByID(id int) (models.User, error)
	CheckInUser(orderId string, userId int) (models.CheckIn, error)
	CheckOutUser(checkInID int) (models.CheckIn, error)
}
