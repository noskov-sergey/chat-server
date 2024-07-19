package model

type UserChat struct {
	ChatID   int    `db:"chat_id"`
	Username string `db:"username"`
}

type Users struct {
	Usernames []string
}
