package router

type UserRequest struct {
	FirstName string `json:"fname" validate:"required"`
	// LastName  string `json:"lname" validate:"required"`
	// Age       uint8  `json:"age" validate:"gte=0,lte=130"`
	// Email     string `json:"email" validate:"required,email"`
}
