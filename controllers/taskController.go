package task

import (
	"Tasklify/Types"
	util "Tasklify/Utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var Tasks = []Types.Task{
	{1, "Homework", "wait"},
	{2, "reading", "done"},
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, Tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetOneById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	validId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		panic(err)
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

func EditTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	taskId, err := strconv.Atoi(id)
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

	for i := range Tasks {
		if Tasks[i].Id == taskId {
			edited := Types.Task{Id: Tasks[i].Id, Title: Tasks[i].Title, Status: Tasks[i].Status}
			if reqBody.Title != "" {
				edited.Title = reqBody.Title
			}
			if reqBody.Status != "" {
				edited.Status = reqBody.Status
			}
			newTasks := util.RemoveElementFromSlice(Tasks, i)
			Tasks = newTasks
			Tasks = append(Tasks, edited)
			break
		}

	}
	fmt.Fprint(w, Tasks)
	return
}
