package domain

type CreateChatWithParticipantsRequestType struct {
	Name            string              `json:"name"`
	RolePermissions map[string][]string `json:"rolePermissions"`
	UserRoles       map[uint]string     `json:"userRoles"`
}

type Chat struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func NewChat(name string) *Chat {
	return &Chat{
		Name: name,
	}
}
