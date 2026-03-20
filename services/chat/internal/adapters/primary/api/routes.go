package api

import "github.com/gin-gonic/gin"

func (h *handler) RegisterRoutes() *gin.Engine {
	r := gin.New()
	// group := r.Group("/chat")
	return r
}
