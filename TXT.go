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
	go copy1(&wg)
	go copy2(&wg)
	go copy3(&wg)
	go copy4(&wg)
	wg.Wait()
	fmt.Println(time.Since(start))
	os.RemoveAll("./TXT/")
}

func fileCreate() {
	err := os.Mkdir("TXT", os.ModePerm)
	if err != nil {
		return
	}
}

func copy1(wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create("./TXT/1.txt")
	if err != nil {
		return
	}
	for x := 1; x <= 1000; x++ {
		defer file.Close()
		file.WriteString("Atiwan\nPhongam\n25\n")
	}
	for i := 1; i <= 750; i++ {
		sFile, err := os.Open(fmt.Sprintf("./TXT/%v.txt", i))
		if err != nil {
			log.Fatal(err)
		}
		defer sFile.Close()

		eFile, err := os.Create(fmt.Sprintf("./TXT/%v.txt", i+1))
		if err != nil {
			log.Fatal(err)
		}
		defer eFile.Close()

		_, err = io.Copy(eFile, sFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
