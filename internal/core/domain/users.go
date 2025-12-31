package domain

import (
	"time"
)

type Users struct {
	BaseModel

	DisplayName    string    `json:"display_name" gorm:"type:varchar(255);not null"`
	FirstName      string    `json:"first_name" gorm:"type:varchar(255);index"`
	LastName       string    `json:"last_name" gorm:"type:varchar(255)"`
	ImageUrl       string    `json:"image_url" gorm:"type:varchar(255)"`
	Username       string    `json:"username" gorm:"type:varchar(50);not null;index"`
	Password       string    `json:"password" gorm:"type:varchar(255);index"`
	Email          string    `json:"email" gorm:"type:varchar(50);not null;index"`
	MobileNumber   string    `json:"mobile_number" gorm:"type:varchar(255);not null"`
	DialCode       string    `json:"dial_code" gorm:"type:varchar(255);not null"`
	ProviderSocial string    `json:"provider_social" gorm:"type:varchar(255)"`
	SocialId       string    `json:"social_id" gorm:"type:varchar(255);index"`
	LastLoginAt    time.Time `json:"last_login_at"`
}

func (Users) TableName() string {
	return "users"
}

func init() {
	RegisterModel(&Users{})
}
