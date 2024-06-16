package model

import (
	"database/sql"

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
