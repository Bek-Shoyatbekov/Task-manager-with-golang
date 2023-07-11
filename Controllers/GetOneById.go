package controller

import (
	util "Tasklify/Utils"
	"fmt"
	"net/http"
)

func GetOneById(w http.ResponseWriter, r *http.Request) {
	validId, err := util.GetAndParseIdFromQueryParam(w, r)
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		panic(err)
		return
	}
	for i := range Tasks {
		if Tasks[i].Id == validId {
			fmt.Fprint(w, Tasks[i])
			return
		}
	}
	fmt.Fprint(w, "Task not found")
	return
}
