package controller

import (
	"Tasklify/Types"
	util "Tasklify/Utils"
	"fmt"
	"net/http"
)

var Tasks = []Types.Task{
	{1, "Homework", "wait"},
	{2, "reading", "done"},
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	taskId, err := util.GetIdFromQueryParam(w, r)
	if err != nil {
		fmt.Fprint(w, "Id is invalid")
		return
	}
	Tasks = util.DeleteElementById(Tasks, taskId)
	fmt.Fprint(w, Tasks)
	return
}
