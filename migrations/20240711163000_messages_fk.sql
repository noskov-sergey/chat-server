-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages
    ADD CONSTRAINT chat_id_fk FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE chat_user
    chat_id integer not null;
-- +goose StatementEnd
