package v1handler

func (h *Handler) MapRoutes() {
	groupHandler := h.server.Group(h.cfg.App.BasePath).Group("/affiliated-dealers")

	groupHandler.POST("/", h.Create)
	groupHandler.GET("/", h.Find)
	groupHandler.GET("/:id", h.FindByID)
	groupHandler.PUT("/:id", h.Update)
}
