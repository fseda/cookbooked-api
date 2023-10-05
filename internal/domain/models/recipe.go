package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Title             string             `gorm:"column:title; size:255; not null; uniqueIndex:unique_user_id_title;" json:"title"`
	Description       string             `gorm:"column:description; not null;" json:"description"`
	Body              string             `gorm:"column:body; not null;" json:"body"`
	Link              string             `gorm:"column:link; size:500; not null;" json:"link"`
	UserID            *uint               `gorm:"column:user_id; uniqueIndex:unique_user_id_title;" json:"user_id"`
	User              User               `gorm:"foreignKey:UserID; references:ID"`
	RecipeTags        []RecipeTag        `gorm:"constraint:OnDelete:CASCADE"`
	RecipeIngredients []RecipeIngredient `gorm:"constraint:OnDelete:CASCADE"`
}

type RecipeTag struct {
	gorm.Model
	ID       uint   `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	RecipeID uint   `gorm:"column:recipe_id; uniqueIndex:recipe_tags_must_be_unique;" json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeID; references:ID"`
	TagID    uint   `gorm:"column:tag_id; uniqueIndex:recipe_tags_must_be_unique;" json:"tag_id"`
	Tag      Tag    `gorm:"foreignKey:TagID; references:ID"`
}

type RecipeIngredient struct {
	gorm.Model
	ID           uint       `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	Quantity     float32    `gorm:"column:quantity; not null;" json:"quantity"`
	RecipeID     uint       `gorm:"column:recipe_id; uniqueIndex:recipe_ingredients_must_be_unique;" json:"recipe_id"`
	Recipe       Recipe     `gorm:"foreignKey:RecipeID; references:ID"`
	IngredientID uint       `gorm:"column:ingredient_id; uniqueIndex:recipe_ingredients_must_be_unique;" json:"ingredient_id"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID; references:ID"`
	UnitID       uint       `gorm:"column:unit_id;" json:"unit_id"`
	Unit         Unit       `gorm:"foreignKey:UnitID; references:ID"`
}
