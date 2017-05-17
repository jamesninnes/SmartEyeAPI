package todoAPI

import (

	_"github.com/lib/pq"
	"database/sql"
)

func upload (t Todo) {

	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteye-testdb" +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com" +
		"password=S!lly5tr1ng")

	if err != nil {
		panic(err)
	}

	db.Ping()

	db.Query("INSERT INTO (name, completed, due) VALUE($1, $2, $3), t.Name, t.Id, t.Due")

}