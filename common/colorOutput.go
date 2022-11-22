package common

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

// CONSOLE COLOR SETUP
var ConsoleBold = color.New(color.Bold).SprintFunc()
var ConsoleSuccess = color.New(color.Bold, color.FgWhite, color.BgGreen).SprintFunc()
var ConsoleInfo = color.New(color.Bold, color.FgWhite, color.BgCyan).SprintFunc()
var ConsoleWarn = color.New(color.Bold, color.FgWhite, color.BgMagenta).SprintFunc()
var ConsoleFailure = color.New(color.Bold, color.FgWhite, color.BgRed).SprintFunc()

func PrettyPrintJSONToConsole(b []byte) {
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
}
