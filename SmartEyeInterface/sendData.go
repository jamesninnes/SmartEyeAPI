package SmartEyeInterface

import (

	_"github.com/lib/pq"
	"database/sql"
)

func UploadTodo (todo Todo)  {

	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteyeTest " +
		"password=S!lly5tr1ng " +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com")

	if err != nil {
		panic(err)
	}

	db.Ping()

	db.Query("INSERT INTO todo (name, completed, due) VALUE($1, $2, $3), t.Name, t.Completed, t.Due")

}


func UploadManualPattern (currPattern manualPattern) errorPayload{

	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteyeTest " +
		"password=S!lly5tr1ng " +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com")

	if err != nil {
		return throwError(1, "Cannot connect to database")
	}

	db.Ping()

	rows, err := db.Query("(SELECT pattern FROM patternTable WHERE pattern = $1), currPattern.pattern")

	if err != nil{
		return throwError(3, "Failed to upload to database. Pattern Already Exists")

	}

	if rows != nil{
		return throwError(2, "Pattern already exists in database")
	}

	db.Query("INSERT INTO patterns (name, pattern, completed,) VALUE($1, $2, $3), pattern.name, pattern.pattern, pattern.created")

	return  throwError(0, "NO ERRORS")

}