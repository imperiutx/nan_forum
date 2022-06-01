package appctx

import "database/sql"

// AppContext app ctx
type AppContext interface {
	GetDBConnection() *sql.DB
}

type appContext struct {
	db *sql.DB
}

func NewAppContext(db *sql.DB) *appContext {
	return &appContext{
		db: db,
	}
}

func (ctx *appContext) GetDBConnection() *sql.DB {
	return ctx.db
}
