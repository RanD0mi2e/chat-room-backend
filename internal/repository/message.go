package repository

type messageRepo struct {
	*Repository
}

type MessageRepo interface {
	GetMessageLimit()
}
