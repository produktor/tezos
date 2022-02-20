-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS telegram_groups
(
    id BIGINT NOT NULL CONSTRAINT telegram_groups_pk PRIMARY KEY,
    title       TEXT,
    description TEXT,
    price       NUMERIC
);

# ALTER TABLE telegram_groups OWNER TO postgres;

CREATE UNIQUE INDEX IF NOT EXISTS telegram_groups_id_uindex ON telegram_groups (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE telegram_groups;
-- +goose StatementEnd
