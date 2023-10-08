package common_test

import (
	"encoding/json"
	"fmt"

	"github.com/mwiater/golangdocker/common"
)

func ExamplePrettyPrintJSONToConsole() {
	testSlice := []string{"one", "one", "two", "two", "three", "three"}
	testSlice1 := common.UniqueSlice(testSlice)
	testSlice2, _ := json.Marshal(testSlice1)
	common.PrettyPrintJSONToConsole(testSlice2)
	// Output:
	// [
	// 	"one",
	// 	"two",
	// 	"three"
	// ]
}

func ExampleUniqueSlice() {
	testSlice := []string{"one", "one", "two", "two", "three", "three"}
	fmt.Println(common.UniqueSlice(testSlice))
	// Output:
	// [one two three]
}

func ExampleSplitStringLines() {
	testString := "line 01\nline 02\nline 03"
	fmt.Println(common.SplitStringLines(testString))
	// Output:
	// [line 01 line 02 line 03]
}
