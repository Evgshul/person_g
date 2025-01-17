package entity

type Person struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id" `
	Fullname    string `gorm:"type:varchar(100)" json:"fullName"`
	Gender      string `gorm:"type:varchar(10)" json:"gender"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phoneNumber"`
	Email       string `gorm:"unique;not null" json:"email"`
}
