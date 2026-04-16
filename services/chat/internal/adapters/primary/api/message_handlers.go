package api

import (
	"github.com/gin-gonic/gin"
)

var NewMessageRequest struct {
	ChatID  uint   `json:"chatID"`
	Content string `json:"string"`
}

func (h *handler) newMessage(ctx *gin.Context) {
	// get the chat receptients by chatID
}

// if err := ctx.ShouldBind(&NewMessageRequest); err != nil {
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": err.Error(),
// 	})
// 	return
// }
//
// currentUserID, ok := ctx.MustGet("userID").(string)
// if !ok {
// 	panic("user id missing")
// }
// currentUserIDint, err := strconv.Atoi(currentUserID)
// if err != nil {
// 	panic("invalid user id type")
// }
// err = h.csvc.NewMessage(uint(currentUserIDint), NewMessageRequest.ChatID, NewMessageRequest.Content)
// if err != nil {
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": err.Error(),
// 	})
// 	return
// }
//
// ctx.JSON(http.StatusCreated, gin.H{
// 	"message": "message created successfully",
// })
// }
