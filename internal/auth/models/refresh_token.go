package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID 			uuid.UUID 		`gorm:"type:uuid;primaryKey" json:"id"`
	UserID 		uuid.UUID 		`gorm:"type:uuid;not null" json:"user_id"`
	Token	 	string 			`gorm:"index;not null" json:"token"`
	ExpiresAt 	time.Time 		`json:"expired_at"`
	Revoked		bool			`gorm:"default:false"`
	CreatedAt 	time.Time 		
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"`
	User      	User 			`gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}