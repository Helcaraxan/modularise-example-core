package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Helcaraxan/modularise-example-core/numberutils"
	"github.com/Helcaraxan/modularise-example-core/stringutils"
)

const defaultStringLength = 20

func main() {
	if len(os.Args) == 0 || len(os.Args) > 2 || (os.Args[1] == "number" && len(os.Args) == 2) {
		fmt.Fprintln(os.Stderr, "Please specify either 'number' or 'string [integer]'.")
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "string":
		stringLength := defaultStringLength
		if len(os.Args) == 2 {
			var err error
			stringLength, err = strconv.Atoi(os.Args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Second argument following 'string' should be an integer.")
				os.Exit(1)
			}
		}
		stringutils.PrintRandomString(stringLength)
	case "number":
		numberutils.PrintRandomNumber()
	default:
	}
}
