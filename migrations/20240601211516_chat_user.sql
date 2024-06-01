-- +goose Up
CREATE TABLE chat_user (
    chat_id integer NOT NULL,
    username varchar NOT NULL,
    created_at timestamp not null default now()
);

-- +goose Down
DROP TABLE chat_user;
