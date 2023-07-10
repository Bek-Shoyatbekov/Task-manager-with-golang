package routes

import (
	controller "Tasklify/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", controller.GetAll)
	myRouter.HandleFunc("/", controller.GetOneById)
	myRouter.HandleFunc("/add", controller.AddTask)
	myRouter.HandleFunc("/edit", controller.EditTask)
	myRouter.HandleFunc("/delete", controller.DeleteTaskById)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
