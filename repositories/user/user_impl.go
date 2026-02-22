package user

import (
	"errors"

	"github.com/jrmygp/user-management/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	var existingUser models.User

	// Check if the username already exists in the database
	err := r.db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		// Username already exists, return a custom error
		return models.User{}, errors.New("username already taken")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// If there's another unexpected error (not "record not found"), return the error
		return models.User{}, err
	}

	// If no errors and the username doesn't exist, create the new user
	err = r.db.Create(&user).Error
	return user, err
}

func (r *repository) EditUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) CreateCheckIn(checkIn models.CheckIn) (models.CheckIn, error) {
	err := r.db.Create(&checkIn).Error
	return checkIn, err
}

func (r *repository) UpdateCheckIn(checkIn models.CheckIn) (models.CheckIn, error) {
	err := r.db.Save(&checkIn).Error
	return checkIn, err
}

func (r *repository) GetCheckInByID(id int) (models.CheckIn, error) {
	var checkIn models.CheckIn
	err := r.db.First(&checkIn, id).Error
	return checkIn, err
}
