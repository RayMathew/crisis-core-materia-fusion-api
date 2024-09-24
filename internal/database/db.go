package database

import (
	"context"
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

func (s *DB) GetAllMateria() ([]crisiscoremateriafusion.Materia, error) {
	var allMateria []crisiscoremateriafusion.Materia
	rows, err := s.DB.Query("SELECT name, materia_type, grade, display_materia_type, description FROM materia")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var m crisiscoremateriafusion.Materia
		err := rows.Scan(&m.Name, &m.Type, &m.Grade, &m.DisplayType, &m.Description)
		if err != nil {
			return nil, err
		}
		allMateria = append(allMateria, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return allMateria, nil
}
