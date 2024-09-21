package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Id   int
	Name string
}

func fetchRows(db *sql.DB) {
	rows, err := db.Query("Select * from data")
	if err != nil {
		log.Panic("Error occured while fetching data")
	}
	count := 0
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.Id, &data.Name)
		if err != nil {
			log.Panic("Error occured while fetching rows")
		}
		count++
		fmt.Println(data)
	}
	log.Printf("%v rows fetched successfully", count)
}

func insertRow(db *sql.DB, row Data) {
	query := fmt.Sprintf("Insert into data values(%d,%q)", row.Id, row.Name)
	fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		log.Panic("Failed to Insert Data into database")
	}
	log.Println(res.LastInsertId())
}

func main() {
	connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Panic("MySQl Connection Error")
	}
	defer db.Close()

	newRow := Data{Id: 4, Name: "Yash"}
	insertRow(db, newRow)
	fetchRows(db)

}
