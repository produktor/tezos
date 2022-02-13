package repository

import (
	"context"
	"link_api/domain/model"
)

type Storage interface {
	AddTgGroup(ctx context.Context, tgGroup model.TelegramGroup) error
	GetTgGroups(ctx context.Context) ([]model.TelegramGroup, error)
	GetTgGroupByID(ctx context.Context, groupID int64) (*model.TelegramGroup, error)
}