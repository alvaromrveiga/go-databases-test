package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	DB *sql.DB
}

func (sql *SQLiteDB) InitializeDB() error {
	sql.DB.Exec("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);")
	return nil
}

func Connect(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)
}

func (sql *SQLiteDB) FindByIdTestTable(id int) (string, error) {
	rows, err := sql.DB.Query("SELECT * FROM user WHERE id = ? ;", id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var rowId string
	var name string

	if rows.Next() {
		err = rows.Scan(&rowId, &name)
		if err != nil {
			return "", err
		}
		return name, nil
	}

	return "", fmt.Errorf("No user found with id %d", id)
}

func (sql *SQLiteDB) UpdateOneTestTable(id int, name string) error {
	_, err := sql.DB.Exec("UPDATE user SET name = ? WHERE id = ?;", name, id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *SQLiteDB) InsertOneTestTable(name string) error {
	_, err := sql.DB.Exec("INSERT INTO user (name) VALUES (?);", name)
	if err != nil {
		return err
	}

	return nil
}
