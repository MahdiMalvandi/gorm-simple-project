package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) (User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	result := db.Create(user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return *user, nil

}

func GetAllUser(db *gorm.DB) ([]User, error) {
	var users []User

	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, username string, updates map[string]interface{}) (User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, err
	}

	if err := db.Model(&user).Updates(updates).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUserByUsername(db *gorm.DB, username string) (bool, error) {
	result := db.Where("username = ?", username).Delete(&User{})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
