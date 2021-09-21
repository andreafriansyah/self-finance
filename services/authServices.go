package services

import (
	"bobee/database"
	"bobee/models"
	"crypto/md5"
	"encoding/hex"
	"log"
)

// GetMD5Hash encryption for user password
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SignIn(email string, password string) (models.User, error) {
	data := models.User{}
	res := database.Conn.
		Table("users").
		Where("users.is_delete = ? AND users.email = ? And users.password = ?", false, email, GetMD5Hash(password)).
		First(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] authService.SignIn : ", err)
		return models.User{}, err
	}
	return data, nil
}
