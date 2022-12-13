package common_test

import (
	"encoding/json"
	"fmt"

	"github.com/mattwiater/golangdocker/common"
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
