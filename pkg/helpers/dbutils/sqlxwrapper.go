// Package dbutils Хелпер-обёртка для выполнения запросов на базе sqlx и для функций подключения к БД (pgx).
package dbutils

// Хелпер-обёртка для выполнения запросов на базе sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
	"github.com/jinzhu/gorm"

)

type Select struct {
	Distinct   bool
	Columns    []Column
	Expression Expression
}

type Expression interface {
	Build(builder Builder)
}

type Delete struct {
	Modifier string
}

type Update struct {
	Modifier string
	Table    Table
}

type TelegramToken struct {
	ID   int
	Name string
}

// sqlErr Форматирование текстов ошибок.
func sqlErr(err error, query string, args ...any) error {
	return fmt.Errorf(`run query "%s" with args %+v: %w`, query, args, err)
}

// namedQuery Заполнение запросов именованными параметрами.
func namedQuery(query string, arg any) (nq string, args []any, err error) {
	nq, args, err = sqlx.Named(query, arg)
	if err != nil {
		return "", nil, sqlErr(err, query, args...)
	}
	return nq, args, nil
}

// Exec Выполнение запросов с параметрами (неименованные, в виде $1...$n).
func Exec(ctx context.Context, db sqlx.ExecerContext, query string, args ...any) (sql.Result, error) {
	res, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return res, sqlErr(err, query, args...)
	}

	return res, nil
}

// NamedExec Выполнение запросов с именованными параметрами.
func NamedExec(ctx context.Context, db sqlx.ExtContext, query string, arg any) (sql.Result, error) {
	nq, args, err := namedQuery(query, arg)
	if err != nil {
		return nil, err
	}

	return Exec(ctx, db, db.Rebind(nq), args...)
}

// GetMap Выборка по запросу с параметрами (неименованные, в виде $1...$n).
// Возвращаемое значение - map - map[string]any
func GetMap(ctx context.Context, db sqlx.QueryerContext, query string, args ...any) (ret map[string]any, err error) {
	row := db.QueryRowxContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, sqlErr(row.Err(), query, args...)
	}

	ret = map[string]any{}
	if err := row.MapScan(ret); err != nil {
		return nil, sqlErr(err, query, args...)
	}

	return ret, nil
}

// TxFunc Описание типа вложенной функции для выполнения в транзакции.
type TxFunc func(tx *sqlx.Tx) error

// TxRunner Интерфейс для запуска транзакции (sqlx).
type TxRunner interface {
	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
}

// RunTx
//
// Запуск транзакции (в случае ошибки выполнения вложенной функции вызовет откат транзакции).
// Вложенная функция (f TxFunc) должна возвращать ошибку в случае присутствия условий, требущих откат транзакции.
func RunTx(ctx context.Context, db TxRunner, f TxFunc) (err error) {
	var tx *sqlx.Tx

	opts := &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	}
	// Запуск транзакции.
	tx, err = db.BeginTxx(ctx, opts)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	// Откат или коммит транзакции при завершении функции.
	defer func() {
		if err != nil {
			// Откат транзакции, т.к. вернулась ошибка.
			err = multierr.Combine(err, tx.Rollback())
		} else {
			// Коммит транзакции.
			err = tx.Commit()
		}
	}()
	// Выполнение вложенной функции и возврат результата.
	return f(tx)
}

// SELECT
func (s Select) Build(builder Builder) {
	if len(s.Columns) > 0 {
		if s.Distinct {
			builder.WriteString("DISTINCT ")
		}

		for idx, column := range s.Columns {
			if idx > 0 {
				builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
	} else {
		builder.WriteByte('*')
	}
}

// DELETE
func (d Delete) Name() string {
	return "DELETE"
}

func (d Delete) Build(builder Builder) {
	builder.WriteString("DELETE")

	if d.Modifier != "" {
		builder.WriteByte(' ')
		builder.WriteString(d.Modifier)
	}
}

func (d Delete) MergeClause(clause *Clause) {
	clause.Name = ""
	clause.Expression = d
}

// Name update clause name
func (update Update) Name() string {
	return "UPDATE"
}

// UPDATE
func (update Update) Build(builder Builder) {
	if update.Modifier != "" {
		builder.WriteString(update.Modifier)
		builder.WriteByte(' ')
	}

	if update.Table.Name == "" {
		builder.WriteQuoted(currentTable)
	} else {
		builder.WriteQuoted(update.Table)
	}
}

// Create new telegram token
func Token() {
	db, err := gorm.Open("postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto create table based on TelegramToken struct
	db.AutoMigrate(&TelegramToken{})

	// Create new telegram token
	token1 := TelegramToken{Name: "telegram token 1"}
	token2 := TelegramToken{Name: "telegram token 2"}
	db.Create(&token1)
	db.Create(&token2)
	fmt.Println("Telegram tokens created successfully")

	// Count number of telegram tokens
	var count int
	db.Model(&TelegramToken{}).Count(&count)
	fmt.Printf("Total number of telegram tokens: %d\n", count)
}

// MergeClause merge update clause
func (update Update) MergeClause(clause *Clause) {
	if v, ok := clause.Expression.(Update); ok {
		if update.Modifier == "" {
			update.Modifier = v.Modifier
		}
		if update.Table.Name == "" {
			update.Table = v.Table
		}
	}
	clause.Expression = update
}
