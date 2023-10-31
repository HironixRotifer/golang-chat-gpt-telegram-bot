// Package db - Работа с хранилищами (базой данных).
package db

// Работа с хранилищем информации о пользователях.

import (
	"context"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/helpers/dbutils"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// UserDataReportRecordDB - Тип, принимающий структуру записей о расходах.
type UserDataReportRecordDB struct {
	Category string `db:"name"`
	Sum      int64  `db:"sum"`
}

// UserStorage - Тип для хранилища информации о пользователях.
type UserStorage struct {
	db              *sqlx.DB
	defaultCurrency string
	defaultLimits   int64
}

// NewUserStorage - Инициализация хранилища информации о пользователях.
// db - *sqlx.DB - ссылка на подключение к БД.
// defaultCurrency - string - валюта по умолчанию.
// defaultLimits - int64 - бюджет по умолчанию.
func NewUserStorage(db *sqlx.DB, defaultCurrency string, defaultLimits int64) *UserStorage {
	return &UserStorage{
		db:              db,
		defaultCurrency: defaultCurrency,
		defaultLimits:   defaultLimits,
	}
}

// InsertUser Добавление пользователя в базу данных.
func (storage *UserStorage) InsertUser(ctx context.Context, userID int64, userName string) error {
	// Запрос на добавление данных.
	const sqlString = `
		INSERT INTO users (tg_id, name, currency, limits)
			VALUES ($1, $2, $3, $4)
			 ON CONFLICT (tg_id) DO NOTHING;`

	// Выполнение запроса на добавление данных.
	if _, err := dbutils.Exec(ctx, storage.db, sqlString, userID, userName, storage.defaultCurrency, storage.defaultLimits); err != nil {
		return err
	}
	return nil
}

// CheckIfUserExist Проверка существования пользователя в базе данных.
func (storage *UserStorage) CheckIfUserExist(ctx context.Context, userID int64) (bool, error) {
	// Запрос на выборку пользователя.
	const sqlString = `SELECT COUNT(id) AS countusers FROM users WHERE tg_id = $1;`

	// Выполнение запроса на получение данных.
	cnt, err := dbutils.GetMap(ctx, storage.db, sqlString, userID)
	if err != nil {
		return false, err
	}
	// Приведение результата запроса к нужному типу.
	countusers, ok := cnt["countusers"].(int64)
	if !ok {
		return false, errors.New("Ошибка приведения типа результата запроса.")
	}
	if countusers == 0 {
		return false, nil
	}
	return true, nil
}

// CheckIfUserExistAndAdd Проверка существования пользователя в базе данных добавление, если не существует.
func (storage *UserStorage) CheckIfUserExistAndAdd(ctx context.Context, userID int64, userName string) (bool, error) {
	exist, err := storage.CheckIfUserExist(ctx, userID)
	if err != nil {
		return false, err
	}
	if !exist {
		// Добавление пользователя в базу, если не существует.
		err := storage.InsertUser(ctx, userID, userName)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
