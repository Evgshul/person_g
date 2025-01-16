package entity

type Person struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Fullname    string `json:"fullName"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
