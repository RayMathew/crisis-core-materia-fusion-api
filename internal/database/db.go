package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	crisiscoremateriafusion "github.com/RayMathew/crisis-core-materia-fusion-api/internal/crisis-core-materia-fusion"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*sqlx.DB
}

func NewConnection(dsn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", "postgres://"+dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	return &DB{db}, nil
}

// User returns a user for a given id.
func (s *DB) User() (*crisiscoremateriafusion.User, error) {
	var u crisiscoremateriafusion.User
	row := s.DB.QueryRow(`SELECT name, exists FROM test WHERE name = 'Ray'`)
	switch err := row.Scan(&u.Name, &u.Exists); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		break
	default:
		panic(err)
	}
	return &u, nil
}
