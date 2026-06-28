package services

import (
	"ego/models"
	"ego/repositories"
)

// GetAllUsers mengambil semua data user untuk dashboard admin
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// GetUserByID mengambil data user berdasarkan ID (untuk admin)
func GetUserByID(id string) (*models.User, error) {
	return repositories.GetUserByID(id)
}
