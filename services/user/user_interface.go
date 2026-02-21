package user

import (
	"errors"

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
