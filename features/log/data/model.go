package data

import (
	"github.com/agussugiyono/coba/features/mentee/data"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Description string
	UserID      uint
	ClassID     uint
	MenteeID    uint
	Mentee      data.Mentee `gorm:"foreignKey:MenteeID"`
}
