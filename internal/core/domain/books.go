package domain

import (
	"time"
)

type Books struct {
	BaseModel

	Title       string    `json:"title" gorm:"type:varchar(255);not null;index"`
	Description string    `json:"description" gorm:"type:text"`
	ImageUrl    string    `json:"image_url" gorm:"type:text;default:null"`
	ISBN        string    `json:"isbn" gorm:"type:varchar(20);uniqueIndex"`
	Price       uint64    `json:"price" gorm:"type:bigint;not null;index"`
	ReleaseDate time.Time `json:"release_date" gorm:"type:date;not null;index"`

	AuthorId    *uint `json:"author_id" gorm:"index"`
	PublisherId *uint `json:"publisher_id" gorm:"index"`

	Author    *Authors   `json:"author,omitempty" gorm:"foreignKey:AuthorId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Publisher *Publisher `json:"publisher,omitempty" gorm:"foreignKey:PublisherId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Books) TableName() string {
	return "books"
}

func init() {
	RegisterModel(&Books{})
}
