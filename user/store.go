package user

import "gorm.io/gorm"

var md *gorm.DB = model()

func createUser(user *User) (uint, error) {
	result := md.Create(&user)

	return user.ID, result.Error
}

func readUser(ip string) (User, error) {
	var user User
	result := md.First(&user, "IP = ?", ip)

	return user, result.Error
}
