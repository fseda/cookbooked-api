package models

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string       `gorm:"column:username; size:255; unique; not null;" json:"username"`
	Email        string       `gorm:"column:email; size:255; unique; not null;" json:"email"`
	PasswordHash string       `gorm:"column:password_hash; not null;" json:"-"`
	Role         role         `gorm:"column:role; default:'user'; not null;" json:"role"`
	Recipes      []Recipe     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"recipes"`
	Tags         []Tag        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tags"`
	Ingredients  []Ingredient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"ingredients"`
	Units        []Unit       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"units"`
}

type role string

const (
	ADMIN role = "admin"
	USER  role = "user"
)

func (r *role) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*r = role(v)
	case string:
		*r = role(v)
	default:
		return fmt.Errorf("unsupported type for role: %T", value)
	}
	return nil
}

func (r role) Value() (driver.Value, error) {
	return string(r), nil
}
