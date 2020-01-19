package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("HelloWorldHelloWorld")
	p := make([]byte, 10)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}

