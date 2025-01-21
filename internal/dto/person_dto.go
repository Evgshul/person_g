package dto

type (
	PersonDto struct {
		FullName    string `json:"fullname"`
		Gender      string `json:"gender"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email" binding:"email"`
	}
	ResponsePersonDto struct {
		ID          int    `json:"id"`
		FullName    string `json:"fullname"`
		Gender      string `json:"gender"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}
)
