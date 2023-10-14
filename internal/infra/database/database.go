package database

import (
	"database/sql"
	"fmt"

	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/gofiber/fiber/v2/log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BootstrapDB(cfg *config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed opening connection to postgres: %v", err)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Info("💾 Database connection established")

	// err = db.AutoMigrate(
	// 	models.User{},
	// 	models.Ingredient{},
	// 	models.RecipeIngredient{},
	// 	models.Recipe{},
	// 	models.RecipeTag{},
	// 	models.Tag{},
	// 	models.Unit{},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}

func CloseDB(db *gorm.DB) error {

	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		return err
	}

	log.Info("Database connection closed")
	return nil
}
