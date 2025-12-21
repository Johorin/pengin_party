package repository

import (
	"context"
	"fmt"
	"pengin_party/config"
	"pengin_party/pkg/app"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/cockroachdb/errors"
)

const (
	MAX_OPEN_CONNS    = 100
	MAX_IDLE_CONNS    = 5
	CONN_MAX_LIFETIME = 5 * time.Minute
)

var (
	ErrNotEntity = errors.New("エンティティが存在しません")
	ErrNotModel  = errors.New("モデルが存在しません")
	ErrNotID     = errors.New("IDが存在しません")
)

type DBInterface interface {
	GetDB() *gorm.DB
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	Close() error
}

type DB struct {
	db *gorm.DB
}

func (d *DB) GetDB() *gorm.DB {
	return d.db
}

func (d *DB) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return d.db.WithContext(ctx).Transaction(fn)
}

func (d *DB) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func Init() (DBInterface, error) {
	dns := config.DB.DNS()
	c := gorm.Config{}
	if app.IsLocal() {
		c = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(mysql.Open(dns), &c)
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました。: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました。: %w", err)
	}

	sqlDB.SetMaxOpenConns(MAX_OPEN_CONNS)
	sqlDB.SetMaxIdleConns(MAX_IDLE_CONNS)
	sqlDB.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	return &DB{db: db}, nil
}