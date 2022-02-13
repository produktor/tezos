package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"link_api/domain/model"
	"link_api/internal/domain/repository"
	"net/http"

	"link_api/internal/config"
	linkRepository "link_api/internal/repository"
	"log"
)

type TelegramService struct {
	bot    *tgbotapi.BotAPI
	r      repository.Storage
	logger *zap.SugaredLogger
}

func NewTelegram(conf config.Config, dbPool *pgxpool.Pool, logger *zap.SugaredLogger) (*TelegramService, error) {
	bot, err := tgbotapi.NewBotAPI(conf.TelegramBotToken)
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	repo := linkRepository.NewLinkRepo(dbPool)

	telegram := TelegramService{
		bot:    bot,
		r:      &repo,
		logger: logger,
	}

	return &telegram, nil
}

func (ts *TelegramService) SetWebhook(webhookAddress string) error {
	wh, _ := tgbotapi.NewWebhook(webhookAddress + "/telegram")

	_, err := ts.bot.Request(wh)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TelegramService) GetBot() (*tgbotapi.BotAPI, error) {
	return ts.bot, nil
}

func (ts *TelegramService) ListenUpdates() {
	updates := ts.bot.ListenForWebhook("/telegram")
	go http.ListenAndServe("0.0.0.0:3000", nil)

	for update := range updates {
		command := update.Message.Command()

		switch command {
		case "start":
			messageText := `To set price send price <value>`
			message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)

			_, err := ts.bot.Send(message)
			if err != nil {
				fmt.Println(err)
			}
		case "price":
			chatID := update.Message.Chat.ID
			chatTitle := update.Message.Chat.Title
			chatDescription := update.Message.Chat.Description
			chatPrice, err := decimal.NewFromString(update.Message.CommandArguments())
			// @TODO нужны проверки на nil

			//получаем чат и записываем в базу
			err = ts.r.AddTgGroup(context.Background(), model.TelegramGroup{
				ID:          chatID,
				Title:       &chatTitle,
				Description: &chatDescription,
				Price:       chatPrice,
			})
			if err != nil {
				fmt.Println(err)
			}

			messageText := fmt.Sprintf("Price %s was set!", update.Message.CommandArguments())
			message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)

			_, err = ts.bot.Send(message)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}