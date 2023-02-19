package db

import (
	"context"
	"git.sample.ru/sample/internal/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DB struct {
	Client *pgxpool.Pool
}

func Get(ctx context.Context, connStr string) (*DB, error) {
	dbConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Error.Printf("unable to parse db config: %s")
	}

	var dbPool *pgxpool.Pool
	for i := 1; i < 8; i++ {
		logger.Info.Printf("trying to connect to the db server (attempt %d)...\n", i)
		dbPool, err = pgxpool.ConnectConfig(ctx, dbConfig)
		if err == nil {
			break
		}
		logger.Error.Printf("got error: %v\n", err)

		// Sleep a bit before trying again
		time.Sleep(time.Duration(i*i) * time.Second)
	}

	// Stop execution if the database was not initialized
	if dbPool == nil {
		logger.Error.Fatalln("could not connect to the database")
	}

	// Get a connection from the pool and check if the database connection is active and working
	db, err := dbPool.Acquire(ctx)
	if err != nil {
		logger.Error.Fatalf("failed to get connection on startup: %v\n", err)
	}
	if err := db.Conn().Ping(ctx); err != nil {
		log.Fatalln(err)
	}

	// Add the connection back to the pool
	db.Release()

	return &DB{
		dbPool,
	}, nil
}

func (d *DB) Close() {
	d.Client.Close()
}
