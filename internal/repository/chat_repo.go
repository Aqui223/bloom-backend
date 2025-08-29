package repository

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/slipe-fun/skid-backend/internal/domain"
)

type ChatRepo struct {
	db *sqlx.DB
}

func NewChatRepo(db *sqlx.DB) *ChatRepo {
	return &ChatRepo{db: db}
}

func (r *ChatRepo) Create(chat *domain.Chat) (*domain.Chat, error) {
	membersJSON, _ := json.Marshal(chat.Members)

	query := `INSERT INTO chats (members) VALUES ($1) RETURNING id, members`

	var created domain.Chat
	var membersBytes []byte
	err := r.db.QueryRow(query, membersJSON).Scan(&created.ID, &membersBytes)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(membersBytes, &created.Members); err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *ChatRepo) GetById(id int) (*domain.Chat, error) {
	var chat domain.Chat
	var membersJSON []byte

	query := `SELECT id, members FROM chats WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&chat.ID, &membersJSON)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(membersJSON, &chat.Members)

	return &chat, nil
}

func (r *ChatRepo) GetByUserId(id int) ([]*domain.Chat, error) {
	rows, err := r.db.Query(`
		SELECT id, members
		FROM chats
		WHERE EXISTS (
			SELECT 1 FROM jsonb_array_elements(members) AS m
			WHERE (m->>'id')::int = $1
		)
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*domain.Chat
	for rows.Next() {
		var chat domain.Chat
		var membersJSON []byte
		if err := rows.Scan(&chat.ID, &membersJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(membersJSON, &chat.Members); err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}
	return chats, rows.Err()
}

func (r *ChatRepo) GetWithUsers(id int, recipient int) (*domain.Chat, error) {
	var chat domain.Chat
	var membersJSON []byte

	query := `
	SELECT *
	FROM chats
	WHERE EXISTS (
		SELECT 1
		FROM jsonb_array_elements(members) AS m
		WHERE (m->>'id')::int = $1
	)
	AND EXISTS (
		SELECT 1
		FROM jsonb_array_elements(members) AS m
		WHERE (m->>'id')::int = $2
	);
	`
	err := r.db.QueryRow(query, id, recipient).Scan(&chat.ID, &membersJSON)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(membersJSON, &chat.Members)

	return &chat, nil
}
