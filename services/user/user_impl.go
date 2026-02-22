package user

import (
	"errors"
	"time"

	"github.com/jrmygp/user-management/grpcclient"
	"github.com/jrmygp/user-management/models"
	"github.com/jrmygp/user-management/repositories/user"
	"github.com/jrmygp/user-management/requests"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type service struct {
	repository user.UserRepository
}

func NewService(repository user.UserRepository) *service {
	return &service{repository}
}

func (s *service) CreateUser(userForm requests.CreateUserRequest) (models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: userForm.Username,
		Password: string(hashedPassword),
	}

	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}

func (s *service) GetUserByID(ID int) (models.User, error) {
	user, err := s.repository.GetUserByID(ID)

	// Check if the error is record not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Return an empty user and nil error to indicate no user found, but no error
		return models.User{}, nil
	}

	// Return the user and any other errors (e.g., DB connection issues)
	return user, err
}

func (s *service) CheckInUser(orderId string, userId int) (models.CheckIn, error) {
	orderClient, conn, err := grpcclient.NewOrderClient()
	if err != nil {
		return models.CheckIn{}, err
	}
	defer conn.Close()

	orderResp, err := orderClient.GetOrderByMidtransID(orderId)

	if err != nil {
		return models.CheckIn{}, err
	}

	if int(orderResp.UserId) != userId {
		return models.CheckIn{}, errors.New("order does not belong to user")
	}

	if orderResp.Status != "paid" {
		return models.CheckIn{}, errors.New("order not paid")
	}

	checkIn := models.CheckIn{
		OrderBookID: int(orderResp.Id),
		UserID:      userId,
		CheckInAt:   time.Now(),
		Status:      "checked_in",
	}

	newCheckIn, err := s.repository.CreateCheckIn(checkIn)
	if err != nil {
		return models.CheckIn{}, err
	}

	return newCheckIn, nil
}

func (s *service) CheckOutUser(checkInID int) (models.CheckIn, error) {
	checkIn, err := s.repository.GetCheckInByID(checkInID)
	if err != nil {
		return models.CheckIn{}, err
	}

	checkIn.Status = "checked_out"
	checkIn.CheckOutAt = time.Now()

	updatedCheckIn, err := s.repository.UpdateCheckIn(checkIn)
	if err != nil {
		return models.CheckIn{}, err
	}

	return updatedCheckIn, nil
}

func (s *service) EditUser(userForm requests.EditUserRequest) (models.User, error) {
	user, err := s.repository.GetUserByID(userForm.UserId)
	if err != nil {
		return models.User{}, err
	}

	user.Balance += userForm.BalanceDelta

	updatedUser, err := s.repository.EditUser(user)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}
