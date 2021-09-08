package models

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Salt string `json:"salt"`
}

func CreateUser(user *User)(err error){
	err = db.Create(&user).Error
	return
}

func GetUserById(id int)(user User){
	db.Where("id = ?",id).First(&user)
	return
}