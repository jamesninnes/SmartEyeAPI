package main

import (
	"encoding/json"
	"fmt"
	_"io"
	_"io/ioutil"
	"net/http"

	_"github.com/gorilla/mux"
	_"time"
	"github.com/k-goepel/smarteyeapi/SmartEyeInterface"
	"time"
	"database/sql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome! test\n")
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	t := SmartEyeInterface.CreateTodo(1, "test", false, time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}



func pullTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Stored in DB:\n")


	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteyeTest " +
		"password=S!lly5tr1ng " +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com")

	if err != nil {
		panic(err)
	}

	db.Ping()

	rows, err := db.Query("SELECT * FROM todo")

	for rows.Next(){
		id := 0
		name := ""
		completed := false
		due := ""
		//t, err := time.Parse("0000-00-00", due)
		//
		//if err != nil{
		//	panic(err)
		//}

		rows.Scan(&id, &name, &completed, &due)
		fmt.Println(rows)

		output := fmt.Sprintf("Id= $1 Name= $2 Completed= $3", id, name, completed)

		fmt.Fprint(w, output)
	}



	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func patternPage(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w, "New Pattern")

	p := SmartEyeInterface.NewManualPattern("Soon", "%1Jw7$hhgeo")
	err := SmartEyeInterface.UploadManualPattern(p)

	output := fmt.Sprintf("$1, $2", SmartEyeInterface.GetMessage(err), SmartEyeInterface.GetTime(err))

	fmt.Fprint(w, output)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

}

func pullPattern(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Stored in DB:\n")


	db, err := sql.Open("postgres", "user=awspostgres " +
		"dbname=smarteyeTest " +
		"password=S!lly5tr1ng " +
		"host=smarteye-testdb.coteiblw5axo.us-west-2.rds.amazonaws.com")

	if err != nil {
		panic(err)
	}

	db.Ping()

	rows, err := db.Query("SELECT * FROM patterns")

	for rows.Next(){
		id := 0
		name := ""
		pattern := ""
		//t, err := time.Parse("0000-00-00", due)
		//
		//if err != nil{
		//	panic(err)
		//}

		rows.Scan(&id, &name, &pattern)
		fmt.Println(rows)

		output := fmt.Sprintf("Id= $1 Name= $2 Pattern= $3", id, name, pattern)

		fmt.Fprint(w, output)
	}



	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}


