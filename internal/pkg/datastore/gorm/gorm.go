package gorm

import (
	"fmt"

	"github.com/oxiginedev/watchman/internal/pkg/datastore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatastore struct {
	*gorm.DB
}

func New(dsn string) (*GormDatastore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	db.AutoMigrate(&datastore.Site{})

	return &GormDatastore{db}, nil
}