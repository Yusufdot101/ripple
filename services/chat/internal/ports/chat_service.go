package ports

type ChatService interface {
	NewChatWithParticipants(userIDs []uint) (uint, error)
}
