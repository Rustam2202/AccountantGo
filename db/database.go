package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	//	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	name     string
	dataBase *sql.DB
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

const TableName = "accounter"

func (db *Database) CreateDataBase(name string) {
	var err error
	db.dataBase, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.dataBase.Close()

	_, err = db.dataBase.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	_, err = db.dataBase.Exec("USE " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.dataBase.Exec(
		fmt.Sprintf(
			`CREATE TABLE IF NOT EXISTS %s ( 
				id INT,
				income FLOAT,
				spend FLOAT,
				date DATETIME,
				comment TEXT
		)`, TableName))
	if err != nil {
		panic(err)
	}

	//mysqlType := fmt.Sprintf(`CREATE DATABASE %s`, db.name)

	/*
		db.name = name
		db.dataBase, db.err = sql.Open("sqlite3", db.name+".db")
		if db.err != nil {
			panic(db.err)
		}


		sqliteType := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY,
			date TEXT,
			income TEXT,
			spend REAL,
			comment TEXT)`,
			db.name)
		statement, errr := db.dataBase.Prepare(sqliteType)
		if db.err != nil {
			panic(errr)
		}
		statement.Exec()
	*/
}

func (db *Database) GetDataBase() *Database {
	return db
}

func (db *Database) AddIncome(income string, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (income, date) VALUES (%f, '%s')`,
		TableName, income, date)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) AddSpend(spend string, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (spend, date) VALUES ('%s', '%s')`,
		TableName, spend, date)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) ShowRecords(date_from string, date_to string) {
	query := fmt.Sprintf(`SELECT * FROM %s `,
		TableName,
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
