package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name               string             `gorm:"column:name; size:100; not null; uniqueIndex:user_ingredients_must_be_unique;" json:"name"`
	Icon               string             `gorm:"column:icon; size:5;" json:"icon"`
	IsSystemIngredient bool               `gorm:"column:is_system_ingredient; default:false; not null; uniqueIndex:user_ingredients_must_be_unique;" json:"is_system_ingredient"`
	UserID             *uint               `gorm:"column:user_id; uniqueIndex:user_ingredients_must_be_unique;" json:"user_id"`
	RecipeIngredients  []RecipeIngredient `gorm:"constraint:OnDelete:RESTRICT"`
}
