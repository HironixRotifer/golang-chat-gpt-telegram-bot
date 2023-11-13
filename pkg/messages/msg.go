package messages

import (
	"context"
	// types "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/bottypes"
)

// MessageSender Интерфейс для работы с сообщениями.
type MessageSender interface {
	SendMessage(text string, userID int64) error
	// ShowInlineButtons(text string, buttons []types.TgRowButtons, userID int64) error
}

// UserDataStorage Интерфейс для работы с хранилищем данных.
type UserDataStorage interface {
	InsertCategory(ctx context.Context, userID int64, catName string, userName string) error
	GetUserCategory(ctx context.Context, userID int64) ([]string, error)
	GetUserCurrency(ctx context.Context, userID int64) (string, error)
	SetUserCurrency(ctx context.Context, userID int64, currencyName string, userName string) error
	GetUserLimit(ctx context.Context, userID int64) (int64, error)
	SetUserLimit(ctx context.Context, userID int64, limits int64, userName string) error
}

// LRUCache Интерфейс для работы с кэшем отчетов.
type LRUCache interface {
	Add(key string, value any)
	Get(key string) any
}

// Model Модель бота (клиент, хранилище, последние команды пользователя)
type Model struct {
	ctx     context.Context
	storage UserDataStorage // Хранилище пользовательской информации.
}

// New Генерация сущности для хранения клиента ТГ и хранилища пользователей и курсов валют.
func New(ctx context.Context, storage UserDataStorage) *Model {
	return &Model{
		ctx:     ctx,
		storage: storage,
	}
}
