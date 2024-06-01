-- +goose Up
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    username varchar,
    chat_id integer not null,
    text text,
    created_at timestamp not null default now()
);

-- +goose Down
DROP TABLE messages;
