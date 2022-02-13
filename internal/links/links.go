package links

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"link_api/domain/model"
	"link_api/internal/domain/repository"
	linkRepository "link_api/internal/repository"
)

type LinkService struct {
	bot    *tgbotapi.BotAPI
	r      repository.Storage
	logger *zap.SugaredLogger
}

func NewLinkService(bot *tgbotapi.BotAPI, dbPool *pgxpool.Pool, logger *zap.SugaredLogger) (*LinkService, error) {
	repo := linkRepository.NewLinkRepo(dbPool)

	return &LinkService{
		bot:    bot,
		r:      &repo,
		logger: logger,
	}, nil
}

func (s *LinkService) GetGroups(ctx context.Context) ([]model.TelegramGroup, error) {
	groups, err := s.r.GetTgGroups(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *LinkService) GetTgLinkByGroupID(ctx context.Context, groupID int64) (string, error) {
	inviteTgConf := tgbotapi.CreateChatInviteLinkConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: groupID,
		},
		Name:               "",
		ExpireDate:         0,
		MemberLimit:        1,
		CreatesJoinRequest: false,
	}

	invite, err := s.bot.Request(inviteTgConf)
	if err != nil {
		fmt.Println(err)
	}

	data, err := invite.Result.MarshalJSON()

	var inv tgInvite
	err = json.Unmarshal(data, &inv)
	if err != nil {
		return "", err
	}

	return inv.InviteLink, nil
}

func (s *LinkService) GetTgLinksByGroupsIDs(ctx context.Context, groupsIDs []int64) ([]model.TelegramLinks, error) {
	telegramLinks := make([]model.TelegramLinks, 0)

	for _, groupID := range groupsIDs {
		inviteTgConf := tgbotapi.CreateChatInviteLinkConfig{
			ChatConfig: tgbotapi.ChatConfig{
				ChatID: groupID,
			},
			Name:               "",
			ExpireDate:         0,
			MemberLimit:        1,
			CreatesJoinRequest: false,
		}

		invite, err := s.bot.Request(inviteTgConf)
		if err != nil {
			fmt.Println(err)
		}

		data, err := invite.Result.MarshalJSON()

		var inv tgInvite
		err = json.Unmarshal(data, &inv)
		if err != nil {
			return nil, err
		}

		group, err := s.r.GetTgGroupByID(ctx, groupID)
		if err != nil {
			return nil, err
		}

		telegramLinks = append(telegramLinks, model.TelegramLinks{
			TgGroup: group,
			Link:    inv.InviteLink,
		})
	}

	return telegramLinks, nil
}

type tgInvite struct {
	InviteLink string `json:"invite_link"`
}
