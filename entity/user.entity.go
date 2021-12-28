package entity


type User struct {
	Firstname	string	`form:"firstname" validate:"required,min=1,max=5"`
	Lastname	string	`form:"lastname" validate:"required,min=1,max=5"`
	Email	string	`form:"email" validate:"required,email"`
}


type UserJson struct {
	Firstname	string	`json:"firstname" binding:"required"`
	Lastname	string	`json:"lastname"  binding:"required"`
}
