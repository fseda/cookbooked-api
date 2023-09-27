package storage

import (
	// "context"
	"database/sql"
	"fmt"
	"log"

	"github.com/fseda/cookbooked-api/internal/infra/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}
	
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, _ := db.DB()
	return sqlDB.Close()
}
