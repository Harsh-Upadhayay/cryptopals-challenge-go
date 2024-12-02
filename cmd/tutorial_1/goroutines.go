package main

import (
	"fmt"
	"sync"
	"time"
)

// If variables are declared outside function then var/const is necessory.
var dbData = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var results = []string{}
var wg = sync.WaitGroup{}
var m = sync.Mutex{}

func dbCall(i int) {
	delay := 1
	time.Sleep(time.Duration(delay) * time.Second)
	m.Lock()
	save(dbData[i])
	m.Unlock()
	wg.Done()
}

func save(result string) {
	results = append(results, result)
}

func main6() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Print("Time: ", time.Since(t0))
	fmt.Print("Results: ", results)
}
