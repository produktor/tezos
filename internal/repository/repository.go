package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"link_api/domain/model"
)

type LinkRepository struct {
	db *pgxpool.Pool
}

func NewLinkRepo(pool *pgxpool.Pool) LinkRepository {
	return LinkRepository{
		db: pool,
	}
}

func (r *LinkRepository) AddTgGroup(ctx context.Context, tgGroup model.TelegramGroup) error {
	price, ok := tgGroup.Price.Float64()
	if !ok {
		// @TODO надо говорить пользователю что неверный формат цены
		price = 0
	}


	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Insert("telegram_groups").Columns("id","title","description","price").
		Values(tgGroup.ID, tgGroup.Title, tgGroup.Description, price).ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(ctx, query, args...);
		err != nil {
		return err
	}

	return nil
}

func (r *LinkRepository) GetTgGroups(ctx context.Context) ([]model.TelegramGroup, error) {
	groups := make([]model.TelegramGroup, 0)

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, _, err := psql.Select("id", "title", "description", "price").
		From("telegram_groups").ToSql()

	rows, err :=  r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var group model.TelegramGroup

		err := rows.Scan(&group.ID, &group.Title, &group.Description, &group.Price)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (r *LinkRepository) GetTgGroupByID(ctx context.Context, groupID int64) (*model.TelegramGroup, error) {
	var group model.TelegramGroup

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Select("id", "title", "description", "price").
		From("telegram_groups").Where(squirrel.Eq{"id": groupID}).ToSql()

	err = r.db.QueryRow(ctx, query, args...).Scan(&group.ID, &group.Title, &group.Description, &group.Price)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *LinkRepository) GetTgGroupsByPrice(ctx context.Context, price decimal.Decimal) ([]model.TelegramGroup, error) {
	groups := make([]model.TelegramGroup, 0)

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Select("id", "title", "description", "price").
		From("telegram_groups").Where(squirrel.LtOrEq{"price": price}).ToSql()

	rows, err :=  r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var group model.TelegramGroup

		err := rows.Scan(&group.ID, &group.Title, &group.Description, &group.Price)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}