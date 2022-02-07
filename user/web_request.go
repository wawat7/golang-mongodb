package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
	City     string `json:"city" binding:"required"`
	Province string `json:"province" binding:"required"`
}

type GetUserDetail struct {
	Id string `uri:"id"`
}
