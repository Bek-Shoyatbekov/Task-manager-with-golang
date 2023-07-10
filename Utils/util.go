package util

import "Tasklify/Types"

func RemoveElementFromSlice(s []Types.Task, i int) []Types.Task {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
