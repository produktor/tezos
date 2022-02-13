package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	telegramBot "link_api/internal/bot/telegram"
	"link_api/internal/config"
	"link_api/internal/database"
	"link_api/internal/links"
	"link_api/internal/links/http/handler"
	"link_api/internal/migrations"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Link api was started!")

	//контекст
	ctx, cancel := context.WithCancel(context.Background())

	//получаем конфиг
	conf, err := config.Init()
	if err != nil {
		fmt.Println(err)
	}

	//инициализируем б.д.
	dbPool, err := database.NewPgSQL(ctx, conf.PostgresDSN)
	if err != nil {
		fmt.Println(err)
	}

	//Инициализируем и выполняем миграции
	err = migrations.Init(conf)
	if err != nil {
		fmt.Println(err)
	}

	//инициализируем логер
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	defer logger.Sync()

	//Инициализируем телеграм
	telegram, err := telegramBot.NewTelegram(*conf, dbPool, sugar)
	if err != nil {
		log.Println(err)
	}

	//Доступ к telegram bot api
	bot, err := telegram.GetBot()
	if err != nil {
		log.Println(err)
	}

	//Создаем LinkService
	linkService, err := links.NewLinkService(bot, dbPool, sugar)
	if err != nil {
		log.Println(err)
	}

	//Устанавливаем вэбхук
	err = telegram.SetWebhook(conf.TelegramWebhookAddress)
	if err != nil {
		log.Println(err)
	}

	//слушаем телеграм обновления
	go telegram.ListenUpdates()

	//обрабатываем наше api
	handler.New(linkService, sugar)

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	select {
	case <-shutdownSignals:
	case <-ctx.Done():
	}

	cancel()
}