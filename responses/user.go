package responses

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

type CheckInResponse struct {
	ID         int    `json:"id"`
	OrderID    int    `json:"order_id"`
	UserID     int    `json:"user_id"`
	CheckInAt  string `json:"check_in_at"`
	CheckOutAt string `json:"check_out_at"`
	Status     string `json:"status"`
}
