// Package common implements utility functions shared accross the application
package common

// Return slice with duplicate items removed
func UniqueSlice(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}

	return result
}
