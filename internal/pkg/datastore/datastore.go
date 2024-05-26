package datastore

import "errors"

var (
	ErrSiteNotCreated = errors.New("site not created")
	ErrSiteNotFound   = errors.New("site not found")
)
