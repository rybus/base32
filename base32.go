package main

import (
	"bufio"
	"encoding/base32"
	"flag"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var decode = flag.Bool("decode", false, "Decode the string")
	flag.Parse()

	for scanner.Scan() {
		if *decode {
			data, err := base32.StdEncoding.DecodeString(scanner.Text())
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			fmt.Printf("%q\n", data)
		} else {
			data := base32.StdEncoding.EncodeToString(scanner.Bytes())
			fmt.Println(data)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
