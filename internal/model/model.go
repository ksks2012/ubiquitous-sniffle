package model

import (
	"database/sql"

	"github.com/ksks2012/ubiquitous-sniffle/global"
	"github.com/ksks2012/ubiquitous-sniffle/pkg/setting"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE IF NOT EXISTS mapping (
  id INTEGER NOT NULL PRIMARY KEY,
  url TEXT NULL,
  short_url TEXT NULL
  );`

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.Path)
	if err != nil {
		return nil, err
	}

	// Check if the table exists
	_, err = db.Exec(create)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveURL(url string, shortURL string) error {
	// Save the URL and shortened URL in the database
	_, err := global.DBEngine.Exec("INSERT INTO mapping (url, short_url) VALUES (?, ?)", url, shortURL)
	return err
}

func GetURL(shortURL string) (string, error) {
	var url string
	err := global.DBEngine.QueryRow("SELECT url FROM mapping WHERE short_url = ?", shortURL).Scan(&url)
	return url, err
}
