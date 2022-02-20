package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"link_api/domain/model"
	"link_api/internal/config"
	"link_api/internal/domain/repository"
	linkRepository "link_api/internal/repository"
	"log"
	"net/http"
	"strings"
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
	go http.ListenAndServe("0.0.0.0:80", nil)

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
		case "value":
			admins, err := ts.bot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
				ChatConfig: tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID},
			})
			if err != nil {
				fmt.Println(err)
			}

			if !isAdmin(update.Message.From.ID, admins) {
				continue
			}

			tgArgument := update.Message.CommandArguments()
			args := strings.Split(tgArgument, " ")

			if len(args) == 0 {
				// @TODO возвращаем сообщение что неверно заданны аргументы, может вообще стоит их тут валидировать
			}

			var newTgGroup model.TelegramGroup
			newTgGroup.ID = update.Message.Chat.ID
			newTgGroup.Title = &update.Message.Chat.Title
			newTgGroup.Description = &update.Message.Chat.Description

			switch args[0] {
			case "ownership":
				newTgGroup.CriteriaType = "ownership"
				newTgGroup.CriteriaToken = &args[1]
			case "balance":
				newTgGroup.CriteriaType = "balance"
				newTgGroup.CriteriaCurrency = &args[1]

				price, err := decimal.NewFromString(args[2])
				if err != nil {
					fmt.Println(err)
				}

				newTgGroup.CriteriaPrice = price
			}

			tgGroup, err := ts.r.GetTgGroups(context.Background())
			if err != nil {
				fmt.Println(err)
			}

			if tgGroup != nil {
				err = ts.r.UpdateTgGroup(context.Background(), newTgGroup)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = ts.r.AddTgGroup(context.Background(), newTgGroup)
				if err != nil {
					fmt.Println(err)
				}
			}

			messageText := fmt.Sprintf("Criteria %s was set!", update.Message.CommandArguments())
			message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)

			_, err = ts.bot.Send(message)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
func isAdmin(userID int64, admins []tgbotapi.ChatMember) bool {
	for _, admin := range admins {
		if admin.User.ID == userID {
			return true
		}
	}

	return false
}
