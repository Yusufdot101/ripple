package api

import (
	"net/http"

	"github.com/Yusufdot101/ripple/services/chat/internal/adapters/primary/api/context"
	"github.com/Yusufdot101/ripple/services/chat/internal/application/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) getConversations(c *gin.Context) {
	currentUserID := context.UserIDFromContext(c)
	q := c.Query("q")

	chats, err := h.csvc.GetChatsByUserID(currentUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var groups, directChats []*domain.Chat
	var otherUserIDs []uint

	for _, chat := range chats {
		members, err := h.csvc.GetChatParticipants(chat.ID, currentUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(members) > 2 {
			groups = append(groups, chat)
			continue
		}

		for _, m := range members {
			if m.UserID != currentUserID {
				directChats = append(directChats, chat)
				otherUserIDs = append(otherUserIDs, m.UserID)
			}
		}
	}

	// get the users in direct chats
	matchedUsers, err := h.csvc.SearchUsers(q, otherUserIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userMatch := make(map[uint32]bool)
	for _, u := range matchedUsers {
		userMatch[u.Id] = true
	}

	var filteredChats []*domain.Chat
	for i, chat := range directChats {
		otherUserID := otherUserIDs[i]
		if userMatch[uint32(otherUserID)] {
			filteredChats = append(filteredChats, chat)
		}
	}

	contacts, err := h.csvc.GetContacts(currentUserID, otherUserIDs, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Groups":   groups,
		"Chats":    filteredChats,
		"Contacts": contacts,
	})
}
