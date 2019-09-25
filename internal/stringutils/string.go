package stringutils

import (
	"fmt"

	"github.com/Helcaraxan/modularise-example-core/internal/random"
)

func PrintRandomString(stringLenght int) {
	fmt.Println(random.GenerateRandomString(stringLenght))
}
