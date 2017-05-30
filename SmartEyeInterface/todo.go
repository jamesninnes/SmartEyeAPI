package SmartEyeInterface

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"created"`
	Due       time.Time `json:"due"`
}

func CreateTodo(Id int, name string, completed bool, due time.Time) Todo {
	var currTodo Todo

	currTodo.Name = name
	currTodo.Id = Id
	currTodo.Completed = completed
	currTodo.Created = time.Now()
	currTodo.Due = due

	return currTodo
}
