package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name              string             `gorm:"column:name; size:50; not null; uniqueIndex:user_units_must_be_unique;" json:"name"`
	IsSystemUnit      bool               `gorm:"column:is_system_unit; default:false; not null; uniqueIndex:user_units_must_be_unique;" json:"is_system_unit"`
	UserID            *uint              `gorm:"column:user_id; uniqueIndex:user_units_must_be_unique;" json:"user_id"`
	RecipeIngredients []RecipeIngredient `gorm:"constraint:OnDelete:RESTRICT"`
}
