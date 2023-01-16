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
	_, _ = fmt.Fprint(os.Stdout, strings.ToLower(string(data)))
}
