package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
}

type UserModel struct {
	DB *gorm.DB
}

func (u *UserModel) AuthenticateUser(username, password string) (*User, error) {
	var user User

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credential")
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credential")
	}
	return &user, nil
}

func (u *UserModel) GetUserByID(id string) (*User, error) {
	var user User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
