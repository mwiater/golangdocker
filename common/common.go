// Package common implements utility functions shared accross the application
package common

import (
	"bufio"
	"strings"
)

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

// Return slice with each line of a multi-line string, splitting on '\n'
func SplitStringLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
