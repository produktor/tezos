-- +goose Up
-- +goose StatementBegin
ALTER TABLE telegram_groups
    ADD COLUMN criteria_type VARCHAR(50) NOT NULL DEFAULT 'balance',
    ADD COLUMN criteria_token TEXT NULL,
    ADD COLUMN criteria_currency TEXT NULL;

ALTER TABLE telegram_groups RENAME COLUMN price TO criteria_price;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_groups RENAME COLUMN criteria_price TO price;

ALTER TABLE telegram_groups
    DROP COLUMN criteria_type,
    DROP COLUMN criteria_token,
    DROP COLUMN criteria_currency;
-- +goose StatementEnd
