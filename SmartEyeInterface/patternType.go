package SmartEyeInterface

import (
	"time"
	"database/sql"
)

type  patternType struct {

	name      string    `json:"name"`
	created   time.Time `json:"created"`

}

func NewPatternType( name string ) {

	var p patternType
	p.name = name
	p.created = time.Now()
	UploadPatternType(p)
}

func UploadPatternType (p patternType) errorPayload{

	db, err := sql.Open("postgres",
		"user=awspostgres " +
			"dbname=smarteyeTest " +
			"password=S!lly5tr1ng " +
			"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com")

	err = db.Ping()

	if err != nil {
		return throwError(1, "Cannot connect to database")
	}

	db.Query("CREATE TABLE $1 (id int PRIMARY KEY, name char(16), pattern string, created time)", p.name)

	return throwError(0, "No Errors")

}