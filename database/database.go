package database

import "github.com/tieubaoca/telegram-dumb-chatbot/types"

type Database interface {
	SaveMessage(chatId int64, from int64, message string) error
	PaginateMessages(chatId int64, page int, limit int) ([]types.Message, error)
}
