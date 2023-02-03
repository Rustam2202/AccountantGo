package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Name     string
	dataBase *sql.DB
}

type CalculateResult struct {
	Data       [][colNumb]string
	AllIncomes float32
	AllSpends  float32
}

const (
	colNumb       = 5
	TableName     = "accounter"
	sqlDateFormat = "2006-01-02"
)

// Using SQLite
func (db *Database) OpenAndCreateLocalDb() error {
	var err error
	db.dataBase, err = sql.Open("sqlite3", fmt.Sprintf("./%s.db", db.Name))
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.dataBase.Exec(
		fmt.Sprintf(
			`CREATE TABLE IF NOT EXISTS %s ( 
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				income REAL,
				spend REAL,
				date TEXT NOT NULL, 
				comment TEXT
		)`, TableName)) // DATE: yyyy-mm-dd sql-format
	if err != nil {
		return err
	}
	return nil
}

// Using MySQL
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
	db.Name = name
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

func (db *Database) AddIncomeAndSpend(income float32, spend float32, date time.Time,
	commentIncome string, commentSpend string) error {

	combineCommet := commentIncome + "\n" + commentSpend
	query := fmt.Sprintf(`INSERT INTO %s (income, spend, date, comment) VALUES (%f, %f, '%s', '%s')`,
		TableName, income, spend, date.Format(sqlDateFormat), combineCommet)

	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) AddIncome(income float32, date time.Time, comment string) error {
	//if err := db.OpenDataBase(db.Name); err != nil {
	//	return err
	//}

	query := fmt.Sprintf(`INSERT INTO %s (income, date, comment) VALUES (%f, '%s', '%s')`,
		TableName, income, date.Format(sqlDateFormat), comment)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) AddSpend(spend float32, date time.Time, comment string) error {
	//if err := db.OpenDataBase(db.Name); err != nil {
	//	return err
	//}

	query := fmt.Sprintf(`INSERT INTO %s (spend, date, comment) VALUES (%f, '%s', '%s')`,
		TableName, spend, date.Format(sqlDateFormat), comment)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) DropTable() error {
	query := fmt.Sprintf(`DROP TABLE %s`,
		TableName)
	statement, err := db.dataBase.Prepare(query)
	if err != nil {
		return err
	}
	//defer db.dataBase.Close()
	statement.Exec()
	return nil
}

func (db *Database) CalculateRecords(dateFrom, dateTo time.Time) (CalculateResult, error) {
	//	if err := db.OpenDataBase(db.Name); err != nil {
	//		return nil, err
	//	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE date >= '%s' AND date <= '%s' ORDER BY date DESC`,
		TableName, dateFrom.Format(sqlDateFormat), dateTo.Format(sqlDateFormat),
	)
	rows, err := db.dataBase.Query(query)
	if err != nil {
		return CalculateResult{}, err
	}

	result := CalculateResult{}
	var allIncomes float32
	var allSpends float32
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
		record[0] = strconv.Itoa(*id)
		d, err := time.Parse(sqlDateFormat, *date)
		if err != nil {
			panic(err)
		}
		record[1] = time.Time.Format(d, "02.01.2006")
		if income != nil {
			allIncomes += *income
			record[2] = fmt.Sprintf("%0.2f", *income)
		}
		if spend != nil {
			allSpends += *spend
			record[3] = fmt.Sprintf("%0.2f", *spend)
		}
		if comment != nil {
			record[4] = *comment
		}
		result.Data = append(result.Data, record)
	}
	result.AllIncomes = allIncomes
	result.AllSpends = allSpends
	return result, nil
}
