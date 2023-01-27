package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	name     string
	dataBase *sql.DB
}

type Record struct {
	Id      *int
	Date    *string // time.Time
	Income  *float32
	Spend   *float32
	Comment *string
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
	// defer db.dataBase.Close()

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
	db.name = name
}

func (db *Database) GetDataBase() *Database {
	return db
}

func (db *Database) AddIncome(income float32, date string) {
	var err error
	db.dataBase, err = sql.Open("mysql", fmt.Sprintf("root:password@tcp(127.0.0.1:3306)/%s", db.name))
	if err != nil {
		panic(err)
	}

	_, err = db.dataBase.Exec("USE " + db.name)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	query := fmt.Sprintf(
		`INSERT %s(income, date) VALUES (%f, "%s")`,
		TableName, income, date)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) AddSpend(spend float32, date string) {
	query := fmt.Sprintf(`INSERT INTO %s (spend, date) VALUES (%f, '%s')`,
		TableName, spend, date)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (db *Database) CalculateRecords(date_from string, date_to string) []Record {
	var err error
	db.dataBase, err = sql.Open("mysql", fmt.Sprintf("root:password@tcp(127.0.0.1:3306)/%s", db.name))
	if err != nil {
		panic(err)
	}
	_, err = db.dataBase.Exec("USE " + db.name)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	query := fmt.Sprintf(`SELECT * FROM %s.%s`,
		db.name, TableName,
	// date_from, date_to,
	)
	rows, err := db.dataBase.Query(query)
	if err != nil {
		panic(err)
	}
	result := []Record{}
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Id, &record.Income, &record.Spend, &record.Date, &record.Comment)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, record)
	}
	return result
}
