package models

type Recipe struct {
	Base
	Title             string              `gorm:"column:title; size:255; not null; uniqueIndex:unique_user_id_title;" json:"title"`
	Description       string              `gorm:"column:description; type:text; not null;" json:"description"`
	Body              string              `gorm:"column:body; type:text; not null;" json:"body"`
	Link              string              `gorm:"column:link; size:500; not null;" json:"link"`
	UserID            *uint               `gorm:"column:user_id; uniqueIndex:unique_user_id_title;" json:"user_id"`
	RecipeTags        []*RecipeTag        `gorm:"constraint:OnDelete:CASCADE" json:"recipe_tags"`
	RecipeIngredients []*RecipeIngredient `gorm:"constraint:OnDelete:CASCADE" json:"recipe_ingredients,omitempty"`
}

type RecipeTag struct {
	Base
	RecipeID uint   `gorm:"column:recipe_id; uniqueIndex:recipe_tags_must_be_unique;" json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeID; references:ID"`
	TagID    uint   `gorm:"column:tag_id; uniqueIndex:recipe_tags_must_be_unique;" json:"tag_id"`
	Tag      Tag    `gorm:"foreignKey:TagID; references:ID"`
}

type RecipeIngredient struct {
	Base
	Quantity float32 `gorm:"column:quantity; not null;" json:"quantity"`
	RecipeID uint    `gorm:"column:recipe_id; uniqueIndex:recipe_ingredients_must_be_unique;" json:"recipe_id"`
	// Recipe       Recipe     `gorm:"foreignKey:RecipeID; references:ID"`
	IngredientID uint       `gorm:"column:ingredient_id; uniqueIndex:recipe_ingredients_must_be_unique;" json:"ingredient_id"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID; references:ID" json:"ingredient,omitempty"`
	UnitID       uint       `gorm:"column:unit_id;" json:"unit_id"`
	Unit         Unit       `gorm:"foreignKey:UnitID; references:ID" json:"unit,omitempty"`
}
