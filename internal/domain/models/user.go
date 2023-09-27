package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint         `gorm:"column:id; primaryKey; autoIncrement; not null;" json:"id"`
	Username     string       `gorm:"column:username; size:255; unique; not null;" json:"userame"`
	Email        string       `gorm:"column:email; size:255; unique; not null;" json:"email"`
	PasswordHash string       `gorm:"column:password_hash; not null;" json:"passwordhash"`
	Recipes      []Recipe     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"recipes"`
	Tags         []Tag        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tags"`
	Ingredients  []Ingredient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"ingredients"`
	Units        []Unit       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"units"`
}
