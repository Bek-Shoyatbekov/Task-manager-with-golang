package controller

import (
	"fmt"
	"net/http"
	"os"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, Tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
