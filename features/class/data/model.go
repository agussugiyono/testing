package data

import (
	"github.com/agussugiyono/coba/features/user/data"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name   string
	UserID uint      // ID pengguna yang memiliki kelas ini
	User   data.User `gorm:"foreignKey:UserID"` // Class memiliki relasi many-to-one dengan User

}
