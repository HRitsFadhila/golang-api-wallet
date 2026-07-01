package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct{
	Id uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	UserId uuid.UUID `json:"user_id" gorm:"type:char(36);index;not null"`
	User User `gorm:"foreignKey:UserId;references:Id"`
	Balance int64 `json:"balance" gorm:"default:0"`
	Status string `json:"status" gorm:"type:varchar(20);default:'active'"`
	AccountNumber string `json:"account_number" gorm:"type:varchar(20);uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func(a *Account) BeforeCreate(tx *gorm.DB)(err error){
	a.Id = uuid.New()
	return
}