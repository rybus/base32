package main

import (
	"bufio"
	"encoding/base32"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data := base32.StdEncoding.EncodeToString(scanner.Bytes())
		fmt.Printf("%q\n", data)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
