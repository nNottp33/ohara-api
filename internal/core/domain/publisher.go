package domain

type Publisher struct {
	BaseModel

	Name string `json:"name" gorm:"type:varchar(255);not null;index"`

	Books []Books `json:"books,omitempty" gorm:"foreignKey:PublisherId"`
}

func (Publisher) TableName() string {
	return "publisher"
}

func init() {
	RegisterModel(&Publisher{})
}
