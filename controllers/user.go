package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jrmygp/user-management/models"
	"github.com/jrmygp/user-management/requests"
	"github.com/jrmygp/user-management/responses"
	"github.com/jrmygp/user-management/services/user"
)

type UserController struct {
	service user.UserService
}

func NewUserController(service user.UserService) *UserController {
	return &UserController{service}
}

func convertUserResponse(o models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:       o.ID,
		Username: o.Username,
		Balance:  o.Balance,
	}
}

func convertCheckInResponse(o models.CheckIn) responses.CheckInResponse {
	return responses.CheckInResponse{
		ID:         o.ID,
		OrderID:    o.OrderBookID,
		UserID:     o.UserID,
		CheckInAt:  o.CheckInAt.Format("2006-01-02T15:04:05Z"),
		CheckOutAt: o.CheckOutAt.Format("2006-01-02T15:04:05Z"),
		Status:     o.Status,
	}
}

func (h *UserController) CreateUser(c *gin.Context) {
	var userForm requests.CreateUserRequest

	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertUserResponse(user),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) FindUserByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	user, err := h.service.GetUserByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If no user is found, return null
	if user.ID == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertUserResponse(user),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) CheckInUser(c *gin.Context) {
	var checkInForm requests.CheckInRequest

	err := c.ShouldBindJSON(&checkInForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	checkIn, err := h.service.CheckInUser(checkInForm.OrderID, checkInForm.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCheckInResponse(checkIn),
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) CheckOutUser(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	checkIn, err := h.service.CheckOutUser(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertCheckInResponse(checkIn),
	}
	c.JSON(http.StatusOK, webResponse)
}
