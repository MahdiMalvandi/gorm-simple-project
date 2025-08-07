package user

import (
	"gorm-project/apps/blog"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey;default:auto_random()"`
	FirstName string	`gorm:"type:varchar(50)"`
	LastName  string	`gorm:"type:varchar(50)"`
	Username  string `gorm:"type:varchar(100);uniqueIndex"`
	Password  string `gorm:"type:varchar(255)"`
	IsAdmin   bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Blogs []blog.Blog `gorm:"foreignKey:AuthorId"`
}
