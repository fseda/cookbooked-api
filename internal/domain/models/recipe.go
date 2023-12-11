package models

type Recipe struct {
	Base
	Title             string              `gorm:"column:title; size:255; not null; uniqueIndex:unique_user_id_title;" json:"title"`
	Description       string              `gorm:"column:description; type:text; not null;" json:"description"`
	Body              string              `gorm:"column:body; type:text; not null;" json:"body"`
	Link              string              `gorm:"column:link; size:500; not null;" json:"link"`
	UserID            *uint               `gorm:"column:user_id; uniqueIndex:unique_user_id_title;" json:"user_id"`
	RecipeTags        []*RecipeTag        `gorm:"constraint:OnDelete:CASCADE" json:"recipe_tags,omitempty"`
}

type RecipeTag struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	RecipeID uint   `gorm:"column:recipe_id; uniqueIndex:recipe_tags_must_be_unique;" json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeID; references:ID"`
	TagID    uint   `gorm:"column:tag_id; uniqueIndex:recipe_tags_must_be_unique;" json:"tag_id"`
	Tag      Tag    `gorm:"foreignKey:TagID; references:ID"`
}
