package domain

import "user/service/pkg/utils/models"

type User struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
	Phone     string `json:"phone" gorm:"validate:required"`
	Blocked   bool   `json:"blocked" gorm:"validate:required"`
}
type TokenUser struct {
	User         models.UserDetails
	AccessToken  string
	RefreshToken string
}
