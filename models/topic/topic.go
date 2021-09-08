package topic

import (
	"fiber-bbs/models/category"
	"fiber-bbs/models/user"
	"gorm.io/gorm"
)

type Topic struct {
	*gorm.Model
	Title           string `gorm:"string;not null"`
	Body            string `gorm:"string;not null;"`
	UserId          uint64 `gorm:"uint;default:0"`
	CategoryId      uint64 `gorm:"uint;default:0"`
	ReplyCount      uint64 `gorm:"uint;default:0"`
	ViewCount       uint64 `gorm:"uint;default:0"`
	LastReplyUserId uint64 `gorm:"uint;default:0"`
	Order           uint64 `gorm:"uint;default:0"`
	Except          string `gorm:"string;default:null"`
	Slug            string `gorm:"string;default:null"`
	User            *user.User
	Category        *category.Category
}
