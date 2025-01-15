package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func main() {
	db, err := sql.Open(
		"pgx",
		"postgres://postgres:Test1234@localhost:5432?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Creating database:")
	_, err = db.Exec("CREATE DATABASE test")
	if err != nil {
		fmt.Println("Cannot create database - might already exist")
		fmt.Println("Dropping table:")
		_, err = db.Exec("DROP TABLE products")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Creating table:")
	_, err = db.Exec("CREATE TABLE products (id varchar(10), name varchar(100), price float)")
	if err != nil {
		panic(err)
	}

	fmt.Println("Preparing insert stmt:")
	stmt, err := db.Prepare("INSERT INTO products VALUES($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calling insert stmt:")
	res, err := stmt.Exec("P-01", "Pizza Margarita", 7.99)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	defer db.Close()
}
