// Package dbutils Хелпер-обёртка для выполнения запросов на базе sqlx и для функций подключения к БД (pgx).
package dbutils

// Хелпер-обёртка для функций подключения к БД (pgx)

import (
	"bytes"
	"context"
	"fmt"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// pgxLogger Логгер для pgx, реализующий интерфейс Logger пакета pgx.
type pgxLogger struct{}

// Log Функция реализации интерфейса Logger пакета pgx.
func (pl *pgxLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]any) {
	var buffer bytes.Buffer
	buffer.WriteString(msg)
	for k, v := range data {
		buffer.WriteString(fmt.Sprintf(" %s=%+v", k, v))
	}
	switch level {
	case pgx.LogLevelTrace, pgx.LogLevelNone, pgx.LogLevelDebug:
		logger.Debug(buffer.String())
	case pgx.LogLevelInfo:
		logger.Info(buffer.String())
	case pgx.LogLevelWarn:
		logger.Warn(buffer.String())
	case pgx.LogLevelError:
		logger.Error(buffer.String())
	default:
		logger.Debug(buffer.String())
	}
}

// NewDBConnect Инициализация подключения к базе данных по заданным параметрам.
func NewDBConnect(connString string) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		logger.Error("Ошибка парсинга строки подключения", "err", err)
		return nil, err
	}
	connConfig.RuntimeParams["application_name"] = "tg-bot"

	connConfig.Logger = &pgxLogger{}
	connConfig.LogLevel = pgx.LogLevelDebug
	connStr := stdlib.RegisterConnConfig(connConfig)
	dbh, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		logger.Error("Ошибка соединения с БД", "err", err)
		return nil, fmt.Errorf("ошибка: prepare db connection: %w", err)
	}
	return dbh, nil
}

// Interface clause interface
type Interface interface {
	Name() string
	Build(Builder)
	MergeClause(*Clause)
}

// ClauseBuilder clause builder, allows to customize how to build clause
type ClauseBuilder func(Clause, Builder)

type Writer interface {
	WriteByte(byte) error
	WriteString(string) (int, error)
}

// Builder builder interface
type Builder interface {
	Writer
	WriteQuoted(field interface{})
	AddVar(Writer, ...interface{})
	AddError(error) error
}

// Clause
type Clause struct {
	Name                string // WHERE
	BeforeExpression    Expression
	AfterNameExpression Expression
	AfterExpression     Expression
	Expression          Expression
	Builder             ClauseBuilder
}

// Build build clause
func (c Clause) Build(builder Builder) {
	if c.Builder != nil {
		c.Builder(c, builder)
	} else if c.Expression != nil {
		if c.BeforeExpression != nil {
			c.BeforeExpression.Build(builder)
			builder.WriteByte(' ')
		}

		if c.Name != "" {
			builder.WriteString(c.Name)
			builder.WriteByte(' ')
		}

		if c.AfterNameExpression != nil {
			c.AfterNameExpression.Build(builder)
			builder.WriteByte(' ')
		}

		c.Expression.Build(builder)

		if c.AfterExpression != nil {
			builder.WriteByte(' ')
			c.AfterExpression.Build(builder)
		}
	}
}

const (
	PrimaryKey   string = "~~~py~~~" // primary key
	CurrentTable string = "~~~ct~~~" // current table
	Associations string = "~~~as~~~" // associations
)

var (
	currentTable  = Table{Name: CurrentTable}
	PrimaryColumn = Column{Table: CurrentTable, Name: PrimaryKey}
)

// Column quote with name
type Column struct {
	Table string
	Name  string
	Alias string
	Raw   bool
}

// Table quote with name
type Table struct {
	Name  string
	Alias string
	Raw   bool
}
