// database-related functions that are reusable

package util

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(sqliteFile string) error {
	var err error

	DB, err = sql.Open("sqlite3", sqliteFile)
	if err != nil { return err }

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, name TEXT, password TEXT)")
	if err != nil { return err }

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS books (id TEXT PRIMARY KEY, title TEXT, author TEXT, description TEXT)")
	if err != nil { return err }

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS essays (id TEXT PRIMARY KEY, language TEXT, author TEXT, title TEXT, content TEXT, meta TEXT)")
	if err != nil { return err }

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS sessions (id TEXT, token TEXT)")
	if err != nil { return err }

	return DB.Ping()
}
