package data

import (
	"OpenBankingAPI/database"
	"OpenBankingAPI/models"
)

func ById(user *models.User, id string) error {
	return database.DB.Where("id = ?", id).First(user).Error
}

func ByEmail(user *models.User, email string) error {
	return database.DB.Where("email = ?", email).First(user).Error
}

func ByRole(users *[]models.User, role string) error {
	return database.DB.Where("role = ?", role).Find(users).Error
}
