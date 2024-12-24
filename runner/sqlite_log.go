package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteLog struct {
	db *sql.DB
}

func NewSqliteLog(path string) (*SqliteLog, error) {
	db, err := sql.Open("sqlite3", path+"?_busy_timeout=300000")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS logs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            message TEXT
        )
    `)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SqliteLog{
		db: db,
	}, nil
}

func (s *SqliteLog) Write(line string) error {
	stmt, err := s.db.Prepare("INSERT INTO logs (message) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(line)
	return err
}

func (s *SqliteLog) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
