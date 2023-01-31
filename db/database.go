package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	name     string
	dataBase *sql.DB
}

const colNumb = 5

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

const (
	TableName     = "accounter"
	sqlDateFormat = "2006-01-02"
)

func (db *Database) CreateDataBase(name string) error {
	var err error
	db.dataBase, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	defer db.dataBase.Close()

	_, err = db.dataBase.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		return err
	}

	_, err = db.dataBase.Exec("USE " + name)
	if err != nil {
		return err
	}

	_, err = db.dataBase.Exec(
		fmt.Sprintf(
			`CREATE TABLE IF NOT EXISTS %s ( 
				id INT PRIMARY KEY AUTO_INCREMENT,
				income FLOAT,
				spend FLOAT,
				date DATE NOT NULL, 
				comment TEXT
		)`, TableName)) // DATE: yyyy-mm-dd sql-format
	if err != nil {
		return err
	}
	db.name = name
	return nil
}

func (db *Database) OpenDataBase(name string) error {
	var err error
	db.dataBase, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	_, err = db.dataBase.Exec("USE " + name)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetDataBase() *Database {
	return db
}

func (db *Database) AddIncome(income float32, date time.Time) error {
	if err := db.OpenDataBase(db.name); err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (income, date) VALUES (%f, "%s")`,
		TableName, income, date.Format(sqlDateFormat))
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) AddSpend(spend float32, date time.Time) error {
	if err := db.OpenDataBase(db.name); err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (spend, date) VALUES (%f, '%s')`,
		TableName, spend, date.Format(sqlDateFormat))
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) CalculateRecords(dateFrom time.Time, dateTo time.Time) ([][colNumb]string, error) {
	if err := db.OpenDataBase(db.name); err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT * FROM %s.%s WHERE date >= '%s' AND date <= '%s'`,
		db.name, TableName, dateFrom.Format(sqlDateFormat), dateTo.Format(sqlDateFormat),
	)
	rows, err := db.dataBase.Query(query)
	if err != nil {
		return nil, err
	}

	result := [][colNumb]string{}
	for rows.Next() {
		var id *int
		var date *string
		var income, spend *float32
		var comment *string
		err := rows.Scan(&id, &income, &spend, &date, &comment)
		if err != nil {
			fmt.Println(err)
		}

		// [0]=id, [1]=date, [2]=income, [3]=spend, [4]=comment; id and date is NOL NULL
		var record [colNumb]string
		record[0] = strconv.Itoa(*id-1)
		d, err := time.Parse(sqlDateFormat, *date)
		if err != nil {
			panic(err)
		}
		record[1] = time.Time.Format(d, "02.01.2006")
		if income != nil {
			record[2] = fmt.Sprintf("%0.2f", *income)
		}
		if spend != nil {
			record[3] = fmt.Sprintf("%0.2f", *spend)
		}
		if comment != nil {
			record[4] = *comment
		}
		result = append(result, record)
	}
	return result, nil
}
