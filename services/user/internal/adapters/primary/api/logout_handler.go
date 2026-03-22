package api

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) logout(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.tsvc.DeleteTokenByStringAndUse(refreshToken, domain.REFRESH)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domain.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
		}
		ctx.String(statusCode, err.Error())
		return
	}

	ctx.String(http.StatusOK, "logged out successfully")
}
