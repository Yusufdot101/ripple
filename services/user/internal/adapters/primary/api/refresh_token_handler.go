package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) RefreshToken(ctx *gin.Context) {
	// get the refreshToken cookie
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	// get accessToken
	accessToken, err := h.tsvc.RefreshAccessToken(refreshToken)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
