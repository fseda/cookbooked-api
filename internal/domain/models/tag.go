package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID          uint        `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	Name        string      `gorm:"column:name; size:100; not null; uniqueIndex:user_tags_must_be_unique" json:"name"`
	IsSystemTag bool        `gorm:"column:is_system_tag; default:false; not null; uniqueIndex:user_tags_must_be_unique" json:"is_system_tag"`
	UserID      uint        `gorm:"column:user_id; uniqueIndex:user_tags_must_be_unique" json:"user_id"`
	RecipeTags  []RecipeTag `gorm:"constraint:OnDelete:CASCADE"`
}
