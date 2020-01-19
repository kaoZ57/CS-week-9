package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("HelloWorld")
	p := make([]byte, 3)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
