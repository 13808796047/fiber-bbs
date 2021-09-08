package category

import (
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model
	Name        string `gorm:"string;not null;index"`
	Description string `gorm:"string;default:null"`
	PostCount   uint64 `gorm:"unit;default:0"`
}
