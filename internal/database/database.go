package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustOpen(dsn string) *gorm.DB {
	db, err := open(dsn)

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	_, err = sqlDB.Exec("SET TIME ZONE 'UTC'")

	if err != nil {
		panic(err)
	}

	return db
}

func open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn))
}
