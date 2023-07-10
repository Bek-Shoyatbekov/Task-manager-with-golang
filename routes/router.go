package routes

import (
	task "Tasklify/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", task.GetAll)
	myRouter.HandleFunc("/", task.GetOneById)
	myRouter.HandleFunc("/add", task.AddTask)
	myRouter.HandleFunc("/edit", task.EditTask)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
