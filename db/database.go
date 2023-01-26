package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	name     string
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

func (db *Database) CreateDataBase(name string) {
	db.name = name
	db.dataBase, db.err = sql.Open("sqlite3", db.name+".db")
	if db.err != nil {
		panic(db.err)
	}
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		date TEXT,
		income TEXT,
		spend REAL,
		comment TEXT)`,
		db.name)
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
		db.name, date, income)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) AddSpend(spend string, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (date, spend) VALUES ('%s', '%s')`,
		db.name, date, spend)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) ShowRecords(date_from string, date_to string) {
	query := fmt.Sprintf(`SELECT * FROM %s `,
		db.name,
	// date_from, date_to,
	)
	rows, err := db.dataBase.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var date *string
		var income *string
		var spend *string
		var comment *string
		err := rows.Scan(&id, &date, &income, &spend, &comment)
		if err != nil {
			fmt.Println(err)
			//continue
		}
		fmt.Println(id, *date, income, spend)
	}
}
