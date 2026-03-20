package domain

type Chat struct {
	ID      uint
	UserIDs []uint
}

type Message struct {
	ID     uint
	ChatID uint
	UserID uint
}
