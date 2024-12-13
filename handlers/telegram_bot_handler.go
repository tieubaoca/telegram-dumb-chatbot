package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/openai/openai-go"
	"github.com/tieubaoca/telegram-dumb-chatbot/app"
	"github.com/tieubaoca/telegram-dumb-chatbot/utils"
)

var _ TelegramHandler = &telegramHandler{}

type TelegramHandler interface {
	HandleTextMessage() bot.HandlerFunc
}

type telegramHandler struct {
	app *app.App
}

func NewTelegramHandler(config *utils.Config) TelegramHandler {
	return &telegramHandler{}
}

func (h *telegramHandler) HandleTextMessage() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}
		messages, err := h.app.DB.PaginateMessages(update.Message.Chat.ID, 1, 20)
		if err != nil {
			h.app.Logger.Error(err)
			return
		}
		msgStringLength := 0
		aiMessages := []openai.ChatCompletionMessage{}

		for _, message := range messages {

		}

	}
}
