package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
	fmt.Print(strings.ToLower(string(data)))
}
