package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID          uint   `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	Title       string `gorm:"column:title; size:255; not null; uniqueIndex:unique_user_id_title;" json:"title"`
	Description string `gorm:"column:description; not null;" json:"description"`
	Body        string `gorm:"column:body; not null;" json:"body"`
	Link        string `gorm:"column:link; size:500; not null;" json:"link"`
	UserID      uint   `gorm:"column:user_id; uniqueIndex:unique_user_id_title;" json:"user_id"`
	RecipeTags  []RecipeTag
}

type RecipeTag struct {
	gorm.Model
	ID       uint   `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	RecipeID uint   `gorm:"column:recipe_id; uniqueIndex: recipe_tags_must_be_unique;" json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeID; references:ID"`
	TagID    uint   `gorm:"column:tag_id; uniqueIndex: recipe_tags_must_be_unique;" json:"tag_id"`
	Tag      Tag    `gorm:"foreignKey:TagID; references:ID"`
}
