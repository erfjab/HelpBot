package database

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB(DBurl string) error {
	db, err := gorm.Open(sqlite.Open(DBurl), &gorm.Config{
		PrepareStmt: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = db
	return nil
}

func PingDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func AutoMigrate() error {
	return DB.AutoMigrate(&Items{})
}

func GetDB() *gorm.DB {
	return DB
}