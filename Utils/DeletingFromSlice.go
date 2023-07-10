package util

import "Tasklify/Types"

func DeleteElementById(tasks []Types.Task, id int) []Types.Task {
	for i := range tasks {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	return tasks
}

func RemoveElementFromSlice(s []Types.Task, i int) []Types.Task {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
