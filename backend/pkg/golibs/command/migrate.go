package command

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Migrate struct {
	db  *sql.DB
	dir string
}

func NewMigrate(dsn string, schema string, dir string) (*Migrate, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("database not available: %s", err)
	}

	err = db.QueryRow("CREATE SCHEMA IF NOT EXISTS " + schema + ";").Scan()
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("error while creating schema: %s", err)
	}

	goose.SetBaseFS(nil)
	goose.SetTableName(schema + ".migrations")

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("error while setting dialect: %s", err)
	}

	return &Migrate{db, dir}, nil
}

func (m *Migrate) Up() error {
	err := goose.Up(m.db, m.dir)
	if err != nil {
		return fmt.Errorf("error while running migrate: %s", err)
	}

	return nil
}

func (m *Migrate) Create(name string) error {
	err := goose.Create(m.db, m.dir, name, "sql")
	if err != nil {
		return fmt.Errorf("error while running migrate-create: %s", err)
	}

	return nil
}

func (m *Migrate) Status() error {
	err := goose.Status(m.db, m.dir)
	if err != nil {
		return fmt.Errorf("error while running migrate-status: %s", err)
	}

	return nil
}
