package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oxiginedev/watchman/internal/pkg/datastore"
)

type siteHandler struct {
	siteRepository datastore.SiteRepository
}

func NewSiteHandler(siteRepo datastore.SiteRepository) *siteHandler {
	return &siteHandler{siteRepository: siteRepo}
}

type createSiteRequest struct {
	URL string `json:"url"`
}

func (s *siteHandler) ListSites(c echo.Context) error {
	return nil
}

func (s *siteHandler) CreateSite(c echo.Context) error {
	var p createSiteRequest

	if err := c.Bind(&p); err != nil {
		return err
	}

	ctx := c.Request().Context()

	err := s.siteRepository.CreateSite(ctx, &datastore.Site{
		URL: p.URL,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (s *siteHandler) ShowSite(c echo.Context) error {
	return nil
}
