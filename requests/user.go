package requests

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type EditUserRequest struct {
	UserId       int `json:"user_id" binding:"required"`
	BalanceDelta int `json:"balance_delta" binding:"required"`
}

type CheckInRequest struct {
	OrderID string `json:"order_id" binding:"required"`
	UserId  int    `json:"user_id" binding:"required"`
}
