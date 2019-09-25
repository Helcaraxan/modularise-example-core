package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Helcaraxan/modularise-example-core/internal/numberutils"
	"github.com/Helcaraxan/modularise-example-core/internal/stringutils"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "Please specify either 'number' or 'string'.")
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "string":
		stringutils.PrintRandomString()
	case "number":
		numberutils.PrintRandomNumber()
	default:
	}
}
