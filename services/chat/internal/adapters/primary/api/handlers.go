package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/ripple/services/chat/internal/application/core/domain"
	"github.com/gin-gonic/gin"
)

var createChatWithParticipantsRequests struct {
	UserIDs []uint `json:"userIDs"`
}

func (h *handler) NewChatWithParticipants(ctx *gin.Context) {
	if err := ctx.ShouldBind(&createChatWithParticipantsRequests); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	currentUserID, ok := ctx.MustGet("userID").(string)
	if !ok {
		panic("user id missing")
	}
	currentUserIDint, err := strconv.Atoi(currentUserID)
	if err != nil {
		panic("invalid user id type")
	}
	createChatWithParticipantsRequests.UserIDs = append(createChatWithParticipantsRequests.UserIDs, uint(currentUserIDint))
	if len(createChatWithParticipantsRequests.UserIDs) < 2 {
		ctx.String(http.StatusBadRequest, "userIDs cannot be less than 2")
		return
	}
	log.Println("here: ", createChatWithParticipantsRequests.UserIDs)
	chatID, err := h.csvc.NewChatWithParticipants(createChatWithParticipantsRequests.UserIDs)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"chatID": chatID,
	})
}

var GetChatRequest struct {
	UserIDs []uint `json:"userIDs"`
}

func (h *handler) GetByUserIDs(ctx *gin.Context) {
	if err := ctx.ShouldBind(&GetChatRequest); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	currentUserID, ok := ctx.MustGet("userID").(string)
	if !ok {
		panic("user id missing")
	}
	currentUserIDint, err := strconv.Atoi(currentUserID)
	if err != nil {
		panic("invalid user id type")
	}

	GetChatRequest.UserIDs = append(GetChatRequest.UserIDs, uint(currentUserIDint))
	if len(GetChatRequest.UserIDs) < 2 {
		ctx.String(http.StatusBadRequest, "userIDs cannot be less than 2")
		return
	}

	chat, err := h.csvc.GetChatByUserIDs(GetChatRequest.UserIDs)
	log.Println("here: ", err)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"chat": chat,
	})
}
