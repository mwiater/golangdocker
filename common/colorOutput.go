package common

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/fatih/color"
)

// Console color varables
var ConsoleBold = color.New(color.Bold).SprintFunc()
var ConsoleSuccess = color.New(color.Bold, color.FgWhite, color.BgGreen).SprintFunc()
var ConsoleInfo = color.New(color.Bold, color.FgWhite, color.BgCyan).SprintFunc()
var ConsoleWarn = color.New(color.Bold, color.FgWhite, color.BgMagenta).SprintFunc()
var ConsoleFailure = color.New(color.Bold, color.FgWhite, color.BgRed).SprintFunc()

// Pretty print JSON to console for debugging
//
// Only active if debug is true in the .env file:
//
//	DEBUG=true
func PrettyPrintJSONToConsole(b []byte) {
	var out bytes.Buffer

	err := json.Indent(&out, b, "", "\t")
	if err != nil {
		log.Println("Error:", err)
	}

	_, err = out.Write([]byte("\n\n"))
	if err != nil {
		log.Println("Error:", err)
	}

	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		log.Println("Error:", err)
	}
}
