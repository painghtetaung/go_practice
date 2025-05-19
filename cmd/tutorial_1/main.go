package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

// type gasEngine struct {
// 	mpg       uint8
// 	gallons   uint8
// 	ownerInfo owner
// }

// type owner struct {
// 	name string
// }

var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func dbCall(i int) {
	// var delay float32 = rand.Float32() * 2000
	var delay int = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()

	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello, World!")

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nresults: %v", results)
}
