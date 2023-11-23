package models

type Unit struct {
	ID                uint               `gorm:"primaryKey" json:"id"`
	Name              string             `gorm:"column:name; size:50; not null; uniqueIndex:user_units_must_be_unique;" json:"name"`
	Symbol            string             `gorm:"column:symbol; size:10; unique;" json:"symbol"`
	Type              Type               `gorm:"column:type;" json:"type"`
	System            System             `gorm:"column:system;" json:"system"`
	IsSystemUnit      bool               `gorm:"column:is_system_unit; default:false; not null; uniqueIndex:user_units_must_be_unique;" json:"is_system_unit"`
	UserID            *uint              `gorm:"column:user_id; uniqueIndex:user_units_must_be_unique;" json:"user_id"`
	RecipeIngredients []RecipeIngredient `gorm:"constraint:OnDelete:RESTRICT"`
}

type Type string
type System string

const (
	WEIGHT      Type = "weight"
	VOLUME      Type = "volume"
	TEMPERATURE Type = "temperature"

	METRIC    System = "metric"
	FARENHEIT System = "farenheit"
)
