package entity

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Posts []Post `json:"posts"`
}
