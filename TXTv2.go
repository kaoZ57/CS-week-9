package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	start := time.Now()
	fileCreate()
	txtCreate()
	go copy1(&wg)
	go copy2(&wg)
	go copy3(&wg)
	go copy4(&wg)
	wg.Wait()
	fmt.Println(time.Since(start))
	//os.RemoveAll("./TXT/")
}
