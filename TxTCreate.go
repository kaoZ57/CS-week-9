package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	start := time.Now()
	err := os.Mkdir("FileTxt", os.ModePerm)
	if err != nil {
		return
	}
	go TXT1(&wg)
	go TXT2(&wg)
	go TXT3(&wg)
	go TXT4(&wg)
	wg.Wait()
	fmt.Println(time.Since(start))
	os.RemoveAll("./FileTxt/")
}

// TXT1 สร้างไฟล์ 1-750
func TXT1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 750; i++ {
		file, err := os.Create(fmt.Sprintf("./FileTxt/%v.txt", i))
		if err != nil {
			return
		}
		for i := 0; i < 1000; i++ {
			defer file.Close()
			file.WriteString("Atiwan\nPhongam\n25\n")
		}
	}
}

// TXT2 สร้างไฟล์ 751-1500
func TXT2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 751; i <= 1500; i++ {
		file, err := os.Create(fmt.Sprintf("./FileTxt/%v.txt", i))
		if err != nil {
			return
		}
		for i := 0; i < 1000; i++ {
			defer file.Close()
			file.WriteString("Atiwan\nPhongam\n25\n")
		}
	}
}

// TXT3 สร้างไฟล์ 1501-2250
func TXT3(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1501; i <= 2250; i++ {
		file, err := os.Create(fmt.Sprintf("./FileTxt/%v.txt", i))
		if err != nil {
			return
		}
		for i := 0; i < 1000; i++ {
			defer file.Close()
			file.WriteString("Atiwan\nPhongam\n25\n")
		}
	}
}

// TXT4 สร้างไฟล์ 2251-3000
func TXT4(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2251; i <= 3000; i++ {
		file, err := os.Create(fmt.Sprintf("./FileTxt/%v.txt", i))
		if err != nil {
			return
		}
		for i := 0; i < 1000; i++ {
			defer file.Close()
			file.WriteString("Atiwan\nPhongam\n25\n")
		}
	}
}
