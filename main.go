package main
import (

  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm-project/apps/blog"
  "gorm-project/apps/user"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(&user.User{}, &blog.Blog{})

  




}