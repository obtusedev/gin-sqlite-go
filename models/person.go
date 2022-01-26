package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IpAddress string `json:"ip_address"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./people.sqlite3")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func GetPersons(count int) ([]Person, error) {
	rows, err := DB.Query("SELECT id, first_name, last_name, email, ip_address from people LIMIT" +
		strconv.Itoa(count))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	people := make([]Person, 0)

	for rows.Next() {
		singlePerson := Person{}
		err = rows.Scan(&singlePerson.Id, &singlePerson.Id, &singlePerson.FirstName,
			&singlePerson.LastName, &singlePerson.Email, &singlePerson.IpAddress)
		if err != nil {
			return nil, err
		}

		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}
