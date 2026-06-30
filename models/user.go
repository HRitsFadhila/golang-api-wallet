package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	Id uuid.UUID `json:"id" gorm:"primarykey;type:char(36)"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func(u *User) BeforeCreate(tx *gorm.DB)(err error){
	u.Id = uuid.New()
	return
}