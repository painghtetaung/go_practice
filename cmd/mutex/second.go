package second

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
	// items := []string{"apple", "banana", "cherry"}
	// for v := range items {
	// 	go func() {
	// 		fmt.Println(items[v])
	// 	}()
	// }

	// var intNum int64 = 10
	// for i := 0; i < int(intNum); i++ {
	// 	fmt.Println(intNum)
	// }

	// var str string = "Hello, World!"
	// str = str + "!"
	// fmt.Println(str)

	// fmt.Println(len(""))

	// var gasEngine gasEngine = gasEngine{
	// 	mpg:     30,
	// 	gallons: 10,
	// 	ownerInfo: owner{
	// 		name: "John Doe",
	// 	},
	// }

	// fmt.Println("gasEngine: ", gasEngine)

	// var users map[string]int = map[string]int{"John": 20, "Jane": 21}

	// users["Pai"]++

	// fmt.Println("users: ", users)

	// if n, ok := users["Pai"]; ok {
	// 	n++
	// 	users["Pai"] = n
	// 	fmt.Println("n: ", n)

	// } else {
	// 	fmt.Println("n not found")
	// }

	// fmt.Println("name: ", name)
	// fmt.Println("ok: ", ok)
	// if ok {
	// 	fmt.Println("name: ", name)
	// } else {
	// 	fmt.Println("name not found")
	// }

	// fmt.Println("users: ", users)

	// pointer
	// var p *int32 = new(int32)
	// var i int32
	// fmt.Println("p: ", p)
	// fmt.Println("i: ", &i)

	// p = &i

	// i = 10

	// fmt.Println("p: ", *p)
	// fmt.Println("i: ", i)

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nresults: %v", results)
}
