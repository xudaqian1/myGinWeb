package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique_index"`
	Salt     string `json:"salt"`
	Role     string `json:"role"`
}

func CreateUser(data map[string]interface{}) (err error) {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
		Salt:     data["salt"].(string),
		Role:     data["role"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id int) (user User) {
	db.Where("id = ?", id).First(&user)
	return
}

func GetUserByUsername(username string) (user *User, err error) {
	user = new(User)
	if err = db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func CheckLogin(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ? AND password =?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
