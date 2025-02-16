package internal

import (
	"fmt"
	"os"
)

func ExitErr(err error) {
	fmt.Printf("ERROR: %s\n", err)
	os.Exit(1)
}