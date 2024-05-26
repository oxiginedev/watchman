package datastore

import "context"

type SiteRepository interface {
	CreateSite(context.Context, *Site) error
	FindSiteByID(context.Context, string) (*Site, error)
}