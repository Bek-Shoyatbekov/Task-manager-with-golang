package controller

import (
	"Tasklify/Types"
	util "Tasklify/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func EditTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	taskId, err := util.GetAndParseIdFromQueryParam(w, r)
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}

	var reqBody Types.Task
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
		return
	}
	if reqBody.Status == "" && reqBody.Title == "" {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	validId := false
	for i := range Tasks {
		if Tasks[i].Id == taskId {
			validId = true
			edited := Types.Task{Id: Tasks[i].Id, Title: Tasks[i].Title, Status: Tasks[i].Status}
			if reqBody.Title != "" {
				edited.Title = reqBody.Title
			}
			if reqBody.Status != "" {
				edited.Status = reqBody.Status
			}
			newTasks := util.RemoveElementFromSliceByIndex(Tasks, i)
			Tasks = newTasks
			Tasks = append(Tasks, edited)
			break
		}

	}
	if validId {
		fmt.Fprint(w, Tasks)
		return
	} else {
		fmt.Fprint(w, "Invalid Id or request")
		return
	}
}
