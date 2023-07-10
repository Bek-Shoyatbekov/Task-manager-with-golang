package util

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetIdFromQueryParam(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.URL.Query().Get("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return taskId, nil
}
