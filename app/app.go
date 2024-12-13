package app

import (
	"context"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	godotenv "github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/sirupsen/logrus"
	"github.com/tieubaoca/telegram-dumb-chatbot/database"
	"github.com/tieubaoca/telegram-dumb-chatbot/services"
	"github.com/tieubaoca/telegram-dumb-chatbot/utils"
)

type App struct {
	RestAPI      *gin.Engine
	TelegramBot  *bot.Bot
	OpenAIClient *openai.Client
	SdService    services.SDService
	Logger       *logrus.Logger
	DB           database.Database
}

func NewApp(config *utils.Config) *App {
	godotenv.Load()

	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile("logs/"+config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err := os.OpenFile("logs/"+config.ErrorFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	if os.Getenv("LOG_LEVEL") != "" {
		level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
		if err != nil {
			panic(err)
		} else {
			logger.SetLevel(level)
		}
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, errorFile)

	restAPI := gin.Default()

	tbot, err := bot.New(
		os.Getenv("BOT_API_KEY"),
	)

	if err != nil {
		logger.WithError(err).Fatal("Failed to create telegram bot")
	}

	tbot.SetWebhook(context.Background(), &bot.SetWebhookParams{
		URL: config.WebhookURI + "/api/v1/telegram-bot/webhook",
	})

}
