package user_service

import "myGinWeb/models"

type User struct {
	ID int
	Username string
	Email string
	Role string
	Password string
	Salt string
}

func (u *User) Add() error{
	user := map[string]interface{}{
		"username": u.Username,
		"role": u.Role,
		"password":u.Password,
		"email":u.Email,
		"salt":u.Salt,
	}
	if err:= models.CreateUser(user); err !=nil{
		return err
	}
	return nil
}