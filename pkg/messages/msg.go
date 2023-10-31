package messages

import (
	"context"
	"time"

	types "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/bottypes"
)

// MessageSender Интерфейс для работы с сообщениями.
type MessageSender interface {
	SendMessage(text string, userID int64) error
	ShowInlineButtons(text string, buttons []types.TgRowButtons, userID int64) error
}

// UserDataStorage Интерфейс для работы с хранилищем данных.
type UserDataStorage interface {
	InsertUserDataRecord(ctx context.Context, userID int64, rec types.UserDataRecord, userName string, limitPeriod time.Time) (bool, error)
	GetUserDataRecord(ctx context.Context, userID int64, period time.Time) ([]types.UserDataReportRecord, error)
	InsertCategory(ctx context.Context, userID int64, catName string, userName string) error
	GetUserCategory(ctx context.Context, userID int64) ([]string, error)
	GetUserCurrency(ctx context.Context, userID int64) (string, error)
	SetUserCurrency(ctx context.Context, userID int64, currencyName string, userName string) error
	GetUserLimit(ctx context.Context, userID int64) (int64, error)
	SetUserLimit(ctx context.Context, userID int64, limits int64, userName string) error
}

// ExchangeRates Интерфейс для работы с курсами валют.
type ExchangeRates interface {
	ConvertSumFromBaseToCurrency(currencyName string, sum int64) (int64, error)
	ConvertSumFromCurrencyToBase(currencyName string, sum int64) (int64, error)
	GetExchangeRate(currencyName string) (float64, error)
	GetMainCurrency() string
	GetCurrenciesList() []string
}

// LRUCache Интерфейс для работы с кэшем отчетов.
type LRUCache interface {
	Add(key string, value any)
	Get(key string) any
}

// kafkaProducer Интерфейс для отправки сообщений в кафку.
type kafkaProducer interface {
	SendMessage(key string, value string) (partition int32, offset int64, err error)
	GetTopic() string
}

// Model Модель бота (клиент, хранилище, последние команды пользователя)
type Model struct {
	ctx             context.Context
	tgClient        MessageSender    // Клиент.
	storage         UserDataStorage  // Хранилище пользовательской информации.
	currencies      ExchangeRates    // Хранилише курсов валют.
	reportCache     LRUCache         // Хранилише кэша.
	kafkaProducer   kafkaProducer    // Кафка
	lastUserCat     map[int64]string // Последняя выбранная пользователем категория.
	lastUserCommand map[int64]string // Последняя выбранная пользователем команда.
}

// New Генерация сущности для хранения клиента ТГ и хранилища пользователей и курсов валют.
func New(ctx context.Context, tgClient MessageSender, storage UserDataStorage, currencies ExchangeRates, reportCache LRUCache, kafka kafkaProducer) *Model {
	return &Model{
		ctx:             ctx,
		tgClient:        tgClient,
		storage:         storage,
		lastUserCat:     map[int64]string{},
		lastUserCommand: map[int64]string{},
		currencies:      currencies,
		reportCache:     reportCache,
		kafkaProducer:   kafka,
	}
}
