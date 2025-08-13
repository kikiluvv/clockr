package db

import (
	"encoding/json"
	"os"
)

type Config struct {
	PayPeriod     string `json:"pay_period"`
	StartOfPeriod string `json:"start_of_period"`
	EndOfPeriod   string `json:"end_of_period"`
}

type Break struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Session struct {
	Date       string  `json:"date"`
	ClockIn    string  `json:"clock_in"`
	ClockOut   string  `json:"clock_out"`
	Breaks     []Break `json:"breaks"`
	TotalHours float64 `json:"total_hours"`
}

type Database struct {
	Config   Config    `json:"config"`
	Sessions []Session `json:"sessions"`
}

var DBFile = "clockr.json"

func Load() (*Database, error) {
	data, err := os.ReadFile(DBFile)
	if err != nil {
		if os.IsNotExist(err) {
			// init empty db
			return &Database{}, nil
		}
		return nil, err
	}
	var db Database
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return &db, nil
}

func Save(db *Database) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DBFile, data, 0644)
}

func AddSession(session Session) error {
	db, err := Load()
	if err != nil {
		return err
	}
	db.Sessions = append(db.Sessions, session)
	return Save(db)
}
