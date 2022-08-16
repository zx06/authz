package db

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunotel"
	"github.com/zx06/authz/internal/biz"
	"github.com/zx06/authz/internal/conf"
)

func NewPG(data *conf.Data) *bun.DB {
	ctx := context.Background()
	conn := pgdriver.NewConnector(pgdriver.WithDSN(data.PgDSN))
	pgdb := sql.OpenDB(conn)
	db := bun.NewDB(pgdb, pgdialect.New())
	db.AddQueryHook(
		bundebug.NewQueryHook(
			// disable the hook
			bundebug.WithEnabled(false),

			// BUNDEBUG=1 logs failed queries
			// BUNDEBUG=2 logs all queries
			bundebug.FromEnv("BUNDEBUG"),
		),
	)
	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(conn.Config().Database)))
	initModels(ctx, db)
	return db
}

func initModels(ctx context.Context, db *bun.DB) error {
	// register models
	m2mModels := []any{
		(*biz.AuthUserPermission)(nil),
		(*biz.AuthUserGroup)(nil),
		(*biz.AuthGroupPermission)(nil),
	}
	db.RegisterModel(m2mModels...)
	// create tables
	models := []any{
		(*biz.AuthUser)(nil),
		(*biz.AuthGroup)(nil),
		(*biz.AuthPermission)(nil),
		(*biz.AuthUserPermission)(nil),
		(*biz.AuthUserGroup)(nil),
		(*biz.AuthGroupPermission)(nil),
	}
	for _, model := range models {
		if _, err := db.NewCreateTable().IfNotExists().Model(model).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
