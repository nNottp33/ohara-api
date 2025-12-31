package domain

type Authors struct {
	BaseModel

	Name string `json:"name" gorm:"type:varchar(255);not null;index"`

	Books []Books `json:"books,omitempty" gorm:"foreignKey:AuthorId"`
}

func (Authors) TableName() string {
	return "authors"
}

func init() {
	RegisterModel(&Authors{})
}
