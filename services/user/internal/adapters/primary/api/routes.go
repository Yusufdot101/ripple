package api

import "github.com/gin-gonic/gin"

func (h *handler) RegisterRoutes() *gin.Engine {
	r := gin.New()
	group := r.Group("/auth")
	group.GET("/google", h.googleBegin)
	group.GET("/google/callback", h.googleCallback)
	group.GET("/refreshtoken", h.RefreshToken)
	return r
}
