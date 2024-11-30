package database

import (
    "database/sql"
    "log"
)

func CreateTables(db *sql.DB) error {
    tables := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        )`,

        `CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        )`,

        `CREATE TABLE IF NOT EXISTS sessions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        session_id TEXT UNIQUE,
        FOREIGN KEY (user_id) REFERENCES users(id)
        )`,
    }

    for i, table := range tables {
        _, err := db.Exec(table)
        if err != nil {
            log.Printf("Error creating table %d: %v\n", i+1, err)
            log.Printf("SQL: %s\n", table)
            return err
        }
    }

    return nil
}