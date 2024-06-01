-- +goose Up
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE chats;
