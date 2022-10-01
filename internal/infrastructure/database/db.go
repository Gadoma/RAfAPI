package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db     *sql.DB
	ctx    context.Context
	cancel func()

	DSN string

	Now func() time.Time
}

var ErrorNotFound error = errors.New("resource not found")

func NewDB(dsn string) *DB {
	db := &DB{
		DSN: dsn,
		Now: time.Now,
	}
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

func (db *DB) Open() (err error) {
	if db.DSN == "" {
		return fmt.Errorf("DSN is required to connect to the database")
	}

	if db.db, err = sql.Open("sqlite3", db.DSN); err != nil {
		return fmt.Errorf("failed opening the database at %s because of %w", db.DSN, err)
	}

	if _, err := db.db.Exec(`PRAGMA journal_mode = wal;`); err != nil {
		return fmt.Errorf("failed enabling `wal` because of %w", err)
	}

	if _, err := db.db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return fmt.Errorf("failed enabling `foreign keys` because of %w", err)
	}

	return nil
}

func (db *DB) Close() error {
	db.cancel()

	if db.db != nil {
		return db.db.Close()
	}

	return nil
}

type Tx struct {
	*sql.Tx
	db  *DB
	now time.Time
}

type StringTime time.Time

func (n *StringTime) Scan(value interface{}) error {
	if value, ok := value.(string); ok {
		*(*time.Time)(n), _ = time.Parse(time.RFC3339, value)
		return nil
	}
	return fmt.Errorf("StringTime failed scanning value %q to time.Time", value)
}

func (n *StringTime) Value() (driver.Value, error) {
	return (*time.Time)(n).UTC().Format(time.RFC3339), nil
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Tx{
		Tx:  tx,
		db:  db,
		now: db.Now().UTC().Truncate(time.Second),
	}, nil
}
