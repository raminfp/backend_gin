package serilizers

type RegisterRequest struct {
	Firstname string `json:"firstname" form:"firstname" binding:"required,min=1"`
	Lastname  string `json:"lastname" form:"lastname" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required,min=4"`
}
