package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorID" json:"author"`
}
