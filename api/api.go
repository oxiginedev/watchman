package api

import "github.com/labstack/echo/v4"

type Handler struct {
	router *echo.Echo

	siteHandler *siteHandler
}

func NewHandler() *Handler {
	siteH := NewSiteHandler()
}

func (h *Handler) Routes() {
	h.router.GET("/sites", h.siteHandler.ListSites)
	h.router.POST("/sites", h.siteHandler.CreateSite)
	h.router.GET("/sites/:id", h.siteHandler.ShowSite)
}