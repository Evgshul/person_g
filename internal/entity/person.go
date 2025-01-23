package entity

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id" `
	Fullname    string `gorm:"type:varchar(100);not null;uniqueIndex" json:"fullName"`
	Gender      string `gorm:"type:varchar(10)" json:"gender"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	Email       string `gorm:"type:varchar(20);not null;uniqueIndex" json:"email"`
}

func (Person) TableName() string {
	return "person"
}
