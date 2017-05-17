package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

func pullData() {

	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteye-testdb" +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com" +
		"password=S!lly5tr1ng")

	if err != nil {
		panic(err)
	}

	db.Ping()

	rows, err := db.Query("SELECT * FROM smarteye-testdb")

	fmt.Print("Rows: $1", rows)

}
