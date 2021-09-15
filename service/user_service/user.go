package user_service

import (
	"myGinWeb/models"
	"myGinWeb/pkg/utils"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
	Password string
	Salt     string
}

func (u *User) Add() error {
	md5Password := utils.MD5(u.Password + u.Salt)
	user := map[string]interface{}{
		"username": u.Username,
		"role":     u.Role,
		"password": md5Password,
		"email":    u.Email,
		"salt":     u.Salt,
	}

	if err := models.CreateUser(user); err != nil {
		return err
	}
	return nil
}
func (u *User) CheckAuth() (bool, error) {
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		return false, err
	}
	md5Password := utils.MD5(u.Password + user.Salt)
	return models.CheckLogin(u.Username, md5Password)
}
