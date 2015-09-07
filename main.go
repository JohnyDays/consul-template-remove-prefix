package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prefix := os.Args[1]
	value := os.Args[2]
	fmt.Print(strings.TrimPrefix(value, prefix))
}
