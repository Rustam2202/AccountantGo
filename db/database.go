package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dataBase *sql.DB
	err      error
}

type Record struct {
	Id      uint
	Date    time.Time
	Income  float32
	Spend   float32
	Comment string
}

type Records struct {
	Records []Record
}

func (db *Database) CreateDataBase() {
	db.dataBase, db.err = sql.Open("sqlite3", "./tutelka.db")
	if db.err != nil {
		panic(db.err)
	}
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		date TEXT,
		income TEXT,
		spend REAL,
		comment TEXT)`,
		"tutelka")
	statement, errr := db.dataBase.Prepare(query)
	if db.err != nil {
		panic(errr)
	}
	statement.Exec()
}

func (db *Database) GetDataBase() *Database {
	return db
}

func (db *Database) AddIncome(income string, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (date, income) VALUES ('%s', '%s')`,
		"tutelka", date, income)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) AddSpend(spend string, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (date, spend) VALUES ('%s', '%s')`,
		"tutelka", date, spend)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) ShowRecords(date_from string, date_to string) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE date > '%s' AND date < '%s'`,
		"tutelka", date_from, date_to)
	rows, err := db.dataBase.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan()
	}
}
