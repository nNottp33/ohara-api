package database

import (
	"fmt"
	"time"

	"github.com/nNottp33/ohara-api/internal/config/env"
	"github.com/nNottp33/ohara-api/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewConnection() (*DB, error) {
	databaseEnv := env.Get[env.DBConfig]("Db")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", databaseEnv.Host, databaseEnv.User,
		databaseEnv.Password, databaseEnv.DBName, databaseEnv.Port,
	)

	connected, errConnected := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errConnected != nil {
		panic("Failed to connect database")
	}

	db, errDB := connected.DB()
	if errDB != nil {
		panic("Failed to get underlying sql.DB")
	}

	db.SetMaxOpenConns(databaseEnv.MaxConnection)
	db.SetMaxIdleConns(databaseEnv.MaxIdleConnection)
	db.SetConnMaxLifetime(5 * time.Minute)

	errMigrate := domain.AutomigrateAll(connected)
	if errMigrate != nil {
		panic("Can not apply migrations")
	}

	return &DB{connected}, nil
}
