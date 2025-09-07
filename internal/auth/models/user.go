package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID 			uuid.UUID 		`gorm:"type:uuid;primaryKey" json:"id"`
	Username 	string 			`gorm:"uniqueIndex;not null" json:"username"`
	Email 		string 			`gorm:"not null" json:"email"`
	Password 	string 			`gorm:"not null" json:"password"`
	Role 		string 			`gorm:"default:user" json:"role"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}