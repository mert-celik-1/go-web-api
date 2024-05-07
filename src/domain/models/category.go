package models

type Category struct {
	BaseModel
	Name     string `gorm:"size:105;type:string;not null;"`
	Products []Product
}
