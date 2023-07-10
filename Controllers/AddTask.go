package controller

import (
	"Tasklify/Types"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, http.StatusBadRequest)
	}
	var reqBody Types.Task
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
		return
	}
	if reqBody.Status == "" || reqBody.Title == "" {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	Tasks = append(Tasks, Types.Task{Id: rand.Intn(100), Title: reqBody.Title, Status: reqBody.Status})
	fmt.Fprint(w, Tasks)
	return
}
