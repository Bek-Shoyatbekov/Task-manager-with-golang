package Routes

import (
	controller "Tasklify/Controllers"
	util "Tasklify/Utils"
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
	port := util.GetEnv("PORT")
	log.Fatal(http.ListenAndServe(port, myRouter))
}
