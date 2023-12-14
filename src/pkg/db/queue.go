package db

import (
	"database/sql"
	_	"github.com/mattn/go-sqlite3"
)

type DatabaseHandler struct {
	db *sql.DB
}

func NewHandler() (*DatabaseHandler, error) {
	db, err := sql.Open("sqlite3", "./db/count.db")
	if err != nil {
		return nil, err
	}
	return &DatabaseHandler{db: db}, nil
}

func (handler *DatabaseHandler) Close() {
	handler.db.Close()
}

func (handler *DatabaseHandler) isExist(url string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM count WHERE url=?)`
	err := handler.db.QueryRow(query, url).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (handler *DatabaseHandler) insert(url string) error {
	query := `INSERT INTO count(url, count, timestamp) VALUES(?, 1, datetime('now'))`
	_, err := handler.db.Exec(query, url)
	if err != nil {
		return err
	}
	return nil
}

func (handler *DatabaseHandler) update(url string) error {
	query := `UPDATE count SET count=count+1, timestamp=datetime('now') WHERE url=?`
	_, err := handler.db.Exec(query, url)
	if err != nil {
		return err
	}
	return nil
}

func (handler *DatabaseHandler) Count(url string) (int, error) {
	exists, err := handler.isExist(url)
	if err != nil {
		return 0, err
	}
	
	if exists {
		err = handler.update(url)
		if err != nil {
			return 0, err
		}
	} else {
		err = handler.insert(url)
		if err != nil {
			panic(err)
			return 0, err
		}
	}

	query := `SELECT count FROM count WHERE url=?`
	var count int
	err = handler.db.QueryRow(query, url).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}