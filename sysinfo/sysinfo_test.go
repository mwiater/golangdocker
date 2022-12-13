package sysinfo_test

import (
	"github.com/mattwiater/golangdocker/sysinfo"
)

func ExampleTestTZ() {
	sysinfo.TestTZ()
	// Output:
	// [ ✓ SUCCESS ] Successfully loaded: America/Los_Angeles
}

func ExampleTestTLS() {
	sysinfo.TestTLS()
	// Output:
	// [ ✓ SUCCESS ] Successfully established https connection to: https://google.com
}
