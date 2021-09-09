package models

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Salt string `json:"salt"`
	Role string `json:"role"`
}

func CreateUser(data map[string]interface{})(err error){
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email: data["email"].(string),
		Salt: data["salt"].(string),
		Role: data["role"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id int)(user User){
	db.Where("id = ?",id).First(&user)
	return
}