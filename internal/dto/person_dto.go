package dto

type (
	PersonDto struct {
		FullName    string `json:"fullname" binding:"required"`
		Gender      string `json:"gender" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
	}
	ResponsePersonDto struct {
		ID          int    `json:"id"`
		FullName    string `json:"fullname"`
		Gender      string `json:"gender"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}
)
