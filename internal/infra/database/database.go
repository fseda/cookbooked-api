package database

import (
	"database/sql"
	"fmt"
	"os"

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

	var logLevel logger.LogLevel
	if os.Getenv("GO_ENV") == "development" {
		log.Info("Debugger mode enabled")
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed opening connection to postgres: %v", err)
	}

	log.Info("💾 Database connection established")

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
