package models

import (
	"database/sql/driver"
	"fmt"
)

type User struct {
	Base
	Username     string   `gorm:"column:username; size:255; unique; not null;" json:"username"`
	Email        string   `gorm:"column:email; size:255; unique; not null;" json:"email"`
	Name         string   `gorm:"column:name; size:255;" json:"name"`
	Bio          string   `gorm:"column:bio; type:text;" json:"bio"`
	Avatar       string   `gorm:"column:avatar_url; size:255;" json:"avatar_url"`
	Location     string   `gorm:"column:location; size:255;" json:"location"`
	PasswordHash string   `gorm:"column:password_hash; not null;" json:"-"`
	Role         Role     `gorm:"column:role; default:'user'; not null;" json:"role"`
	Recipes      []Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"recipes"`
	Tags         []Tag    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tags"`
	GithubID     string   `gorm:"column:github_id;" json:"github_id"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

func (r *Role) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*r = Role(v)
	case string:
		*r = Role(v)
	default:
		return fmt.Errorf("unsupported type for role: %T", value)
	}
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
