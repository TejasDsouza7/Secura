package db

import (
    "database/sql"
    "log"
    "golang.org/x/crypto/bcrypt"
    _ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "file:data.db?cache=shared&_foreign_keys=on")
    if err != nil {
        log.Fatal(err)
    }

    createUsers := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password_hash TEXT
    );`

    createFiles := `CREATE TABLE IF NOT EXISTS files (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        filename TEXT,
        uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    DB.Exec(createUsers)
    DB.Exec(createFiles)

    var exists int
    DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'Your Username'").Scan(&exists)
    if exists == 0 {
        hash, _ := bcrypt.GenerateFromPassword([]byte("Your Password"), bcrypt.DefaultCost)
        DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", "Your Username", string(hash))
    }
}