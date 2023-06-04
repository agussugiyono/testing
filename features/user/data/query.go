package data

import (
	"fmt"

	"github.com/agussugiyono/coba/features/user"
	"github.com/agussugiyono/coba/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Insert implements user.UserDataInterface
func (repo *userQuery) Insert(user *user.Core) error {
	if user.Role != helper.NewUserRole("admin") && user.Team != helper.NewUserTeam("manager") {
		return fmt.Errorf("only admin and manager can add users")
	}

	// Create a new database model from the user core data
	userData := ModelToCore(user)

	// Hash password sebelum disimpan
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Mengganti password dengan hashed password
	userData.Password = hashedPassword

	// Insert the user data into the database
	if err := repo.db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
