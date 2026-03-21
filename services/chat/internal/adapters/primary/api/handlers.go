package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// will add jwt later to verify the user is actually there as well because you cant be creating a chat that you are not a part of
var createChatWithParticipantsRequests struct {
	UserIDs []uint `json:"userIDs"`
}

func (h *handler) NewChatWithParticipants(ctx *gin.Context) {
	if err := ctx.ShouldBind(&createChatWithParticipantsRequests); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if len(createChatWithParticipantsRequests.UserIDs) < 2 {
		ctx.String(http.StatusBadRequest, "userIDs cannot be less than 2")
		return
	}
	log.Println("here:", createChatWithParticipantsRequests)
	chatID, err := h.csvc.NewChatWithParticipants(createChatWithParticipantsRequests.UserIDs)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"chatID": chatID,
	})
}
