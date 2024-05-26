package gorm

import (
	"context"
	"errors"

	"github.com/oxiginedev/watchman/internal/pkg/datastore"
	"gorm.io/gorm"
)

type siteRepository struct {
	db *gorm.DB
}

func NewSiteRepository(db *gorm.DB) *siteRepository {
	return &siteRepository{db: db}
}

func (s *siteRepository) CreateSite(ctx context.Context, site *datastore.Site) error {
	result := s.db.WithContext(ctx).Create(site)
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return datastore.ErrSiteNotCreated
	}

	return nil
}

func (s *siteRepository) FindSiteByID(ctx context.Context, id string) (*datastore.Site, error) {
	var site datastore.Site

	err := s.db.WithContext(ctx).Where("id = $1", id).First(&site).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, datastore.ErrSiteNotFound
		}
		return nil, err
	}

	return &site, nil
}