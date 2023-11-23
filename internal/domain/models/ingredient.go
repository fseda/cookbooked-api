package models

type Ingredient struct {
	ID                uint               `gorm:"primaryKey" json:"id"`
	Name               string              `gorm:"column:name; size:100; not null; uniqueIndex:user_ingredients_must_be_unique;" json:"name"`
	Icon               string              `gorm:"column:icon; size:5;" json:"icon"`
	IsSystemIngredient bool                `gorm:"column:is_system_ingredient; default:false; not null; uniqueIndex:user_ingredients_must_be_unique;" json:"is_system_ingredient"`
	UserID             *uint               `gorm:"column:user_id; uniqueIndex:user_ingredients_must_be_unique;" json:"user_id"`
	CategoryID         *uint               `gorm:"column:category_id; not null;" json:"category_id"`
	Category           IngredientsCategory `gorm:"foreignKey:CategoryID; references:ID" json:"category,omitempty"`
	RecipeIngredients  []RecipeIngredient  `gorm:"constraint:OnDelete:RESTRICT" json:"recipe_ingredients,omitempty"`
}

type IngredientsCategory struct {
	ID                uint               `gorm:"primaryKey" json:"id"`
	Category    string `gorm:"column:category; size:100; unique; not null;" json:"category"`
	Description string `gorm:"column:description; type:text" json:"description"`
}

func (IngredientsCategory) TableName() string {
	return "ingredients_categories"
}
