package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/isutare412/oauth-gateway/internal/core/model"
	"github.com/isutare412/oauth-gateway/internal/log"
)

type Client struct {
	db *gorm.DB
}

func NewClient(cfg Config) (*Client, error) {
	db, err := gorm.Open(
		postgres.Open(buildDataSourceName(cfg)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			TranslateError:                           true,
			Logger:                                   log.NewGORMLogger(cfg.SlowQueryThreshold),
		})
	if err != nil {
		return nil, fmt.Errorf("opening gorm db: %w", err)
	}

	return &Client{
		db: db,
	}, nil
}

func (c *Client) Initialize(ctx context.Context) error {
	if err := c.db.WithContext(ctx).AutoMigrate(
		&model.APIToken{},
		&model.User{},
		&model.UserApplicationRole{},
		&model.GoogleAccount{},
		&model.Application{},
		&model.AuthorizedOrigin{},
		&model.AuthorizedRedirectionURI{},
	); err != nil {
		return fmt.Errorf("migrating schemas: %w", err)
	}
	return nil
}

func (c *Client) BeginTx(
	ctx context.Context,
	opts ...*sql.TxOptions,
) (ctxWithTx context.Context, commit, rollback func() error) {
	if _, ok := extractTransaction(ctx); ok {
		panic("nested transaction detected")
	}

	tx := c.db.Begin(opts...)
	ctxWithTx = injectTransaction(ctx, tx)

	commit = func() error {
		return tx.Commit().Error
	}

	rollback = func() error {
		return tx.Rollback().Error
	}

	return ctxWithTx, commit, rollback
}

func (c *Client) WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	ctxWithTx, commit, rollback := c.BeginTx(ctx)

	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("panicked during transaction: %v", v)

			if rerr := rollback(); rerr != nil {
				err = errors.Join(err, fmt.Errorf("during transaction rollback after panic recover: %w", rerr))
			}
		}
	}()

	if ferr := fn(ctxWithTx); ferr != nil {
		if rerr := rollback(); rerr != nil {
			ferr = errors.Join(ferr, fmt.Errorf("during transaction rollback: %w", rerr))
		}
		return ferr
	}

	if err := commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func buildDataSourceName(cfg Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
}

type contextKeyTransaction struct{}

func injectTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, contextKeyTransaction{}, tx)
}

func extractTransaction(ctx context.Context) (tx *gorm.DB, ok bool) {
	tx, ok = ctx.Value(contextKeyTransaction{}).(*gorm.DB)
	return tx, ok
}

func getTxOrDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := extractTransaction(ctx); ok {
		return tx
	}
	return db
}
