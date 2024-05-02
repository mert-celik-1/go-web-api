package models

type Product struct {
	Name       string   `gorm:"size:105;type:string;not null;"`
	Price      float64  `gorm:"type:decimal(10,2);not null"`
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CategoryId int
}
