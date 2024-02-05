package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// test()
	// arraySlice()
	// slice()
	// jsonMap()
	// goRoutine()
	goChainChannels()
	// goMultipleChannels()
	// goWaitGroup()
	// goMutex()
	// goAtomicity()

	// goViper()
	// testi()
	// cast()
	// httpCustomMux()
	// httpDefaultMux()
	
	// var ctx context.Context
	// ctx.Deadline()
}

func httpCustomMux() {
	var health = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func httpDefaultMux() {
	var health = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	}
	http.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test() {
	var sli [3]int
	sli[1] = 4
	fmt.Println(sli)
	slices := make([]int, 5)
	slices[0] = 1
	fmt.Println("slices", slices)
}

func arraySlice() {
	fmt.Println("\nARRAY VS SLICE")
	var array = [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	fmt.Println("array before join function", array)
	fmt.Println("slice before join function", slice)
	changeArray(array)
	changeSlice(slice)
	fmt.Println("array after join function", array)
	fmt.Println("slice after join function", slice)

}

func changeArray(arr [3]int) {
	arr[0] = 2
	fmt.Println("array inside function", arr)
}

func changeSlice(slice []int) {
	slice[0] = 2
	fmt.Println("slice inside function", slice[0])
	slice =  append(slice, 3)
}

func slice() {
	fmt.Println("\nSLICE")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	// WHY THE CAPACITY CHANGE AFTER DROP FIRST TWO VALUES?
	// because capacity meaning how many elements max after that first pointer
	// https://www.youtube.com/watch?v=KzKNGGoaT5U 4:50
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func jsonMap() {
	var jsonNil map[string]string
	// In maps, the zero value of the map is nil and a nil map doesn’t contain any key.
	// If you try to add a key-value pair in the nil map, then the compiler will throw runtime error.
	// jsonNil["name"] = "cris" -> error

	if jsonNil == nil {
		fmt.Println("true")
	}

	json := make(map[string]string)
	json["username"] = "cris"
	fmt.Println(json)

	jsonConstructor := map[string]string{
		"username": "cris",
		"lastname": "rumbo",
	}
	fmt.Println(jsonConstructor)

	value, ok := jsonConstructor["username"]
	fmt.Printf("\nvalue %s, ok %t", value, ok)

	for key, val := range jsonConstructor {
		fmt.Printf("key %s, val %s \n", key, val)
	}

	jsonEmpty := map[string]string{}
	jsonEmpty["age"] = "8"
	fmt.Println(jsonEmpty)
}

func goRoutine() {
	go goRoutineExample() // Create a Goroutine
	fmt.Println("Main function")

	time.Sleep(3 * time.Second) // Sleep to allow Goroutine to finish
}

func goRoutineExample() {
	for i := 0; i < 10; i++ {
		fmt.Println("Routine", i)
		time.Sleep(time.Second) // Simulate some work
	}
}

// func reminder1() {
// 	var numberInt int
// 	numberInt = 3
// 	number := 3
// 	var array [3]int
// 	array[0] = 1
// 	array2 := [3]int{1, 2}
// 	var slice []int
// 	slice[0] = 1
// 	slice2 := make([]int, 3)
// 	slice2[0] = 1
// 	slice3 := slice2[:4]
// 	slice4 := slice3[:2]
// 	json := map[string]string{"string": "string"}
// 	json1 := make(map[string] string)
// 	json1["sring"] = "string"

// 	var goru = func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("Routine", i)
// 			time.Sleep(time.Second) // Simulate some work
// 		}
// 	}

// 	go goru();
// 	time.Sleep(3  * time.Second) // Simulate some work
// }

func goChainChannels() {
	dataChannel1 := make(chan int)
	dataChannel2 := make(chan int)

	go goRoutineChannel1(dataChannel1)
	go goRoutineChannel2(dataChannel1, dataChannel2)

	for {
		val2, ok := <-dataChannel2
		if !ok {
			fmt.Println("Break recieve from 2")
			break
		}
		fmt.Println("Recieved from go routine 2, value ", val2)
	}

}

func goRoutineChannel1(ch chan int) {
	for i := 1; i < 5; i++ {
		ch <- i
		// time.Sleep(time.Second)
	}
	fmt.Println("Close routine 1")
	close(ch)
}

func goRoutineChannel2(ch1, ch2 chan int) {
	for {
		val, ok := <-ch1
		if !ok {
			fmt.Println("Break recieve from 1")
			break
		}
		fmt.Println("Recieved from go routine 1, value ", val)
		val2 := val + 1
		ch2 <- val2
	}
	fmt.Println("Close routine 2")
	close(ch2)
}

func goMultipleChannels() {
	data1 := make(chan int)
	data2 := make(chan int)

	goChannel1 := func(ch chan int) {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}

	goChannel2 := func(ch chan int) {
		for i := 5; i < 15; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}

	go goChannel1(data1)
	go goChannel2(data2)

	for {
		select {
		case msg1, ok := <-data1:
			if !ok {
				break
			}
			fmt.Println("Recieved from 1, value ", msg1)
		case msg2, ok := <-data2:
			if !ok {
				break
			}
			fmt.Println("Recieved from 2, value ", msg2)
		}
	}
}

func goWaitGroup() {
	goWaitGroupRoutine := func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go goWaitGroupRoutine(&wg)
	go goWaitGroupRoutine(&wg)
	wg.Wait()
}

func goMutex() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	var counter uint
	goRoutineMutex := func(wg *sync.WaitGroup) {
		mutex.Lock()
		counter = counter + 1
		fmt.Println(counter)
		mutex.Unlock()
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go goRoutineMutex(&wg)
	}
	wg.Wait()
}

func goAtomicity() {
	var counter int64 // Int64 for 64-bit atomic operations

	// Increment counter in a Goroutine
	go func() {
		for i := 0; i < 5; i++ {
			atomic.AddInt64(&counter, 1)
			// counter++
			fmt.Println("+", counter)
			time.Sleep(time.Millisecond)
		}
	}()

	// Decrement counter in another Goroutine
	go func() {
		for i := 0; i < 5; i++ {
			atomic.AddInt64(&counter, -1)
			// counter--
			fmt.Println("-", counter)
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(time.Second) // Allow Goroutines to finish

	// Load and print the final counter value
	finalValue := atomic.LoadInt64(&counter)
	fmt.Println("Final Counter Value:", finalValue)
}

func goViper() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	initializeViper(env)

	fmt.Printf("App Name: %s\n", viper.GetString("app.name"))
	fmt.Printf("DB Host: %s\n", viper.GetString("db.host"))
}

func initializeViper(env string) {
	viper.SetConfigName(env)    // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension
	viper.AddConfigPath(".env") // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Fatalf("Error while reading config file %s", err)
	}
}

func testi() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	// var val int
	for _, v := range a {
		fmt.Println("a", a)
		go func() {
			fmt.Println("v", v)
			ch <- v * 2
		}()
		// val = <- ch
	}
	fmt.Println("vee")
	for i := 0; i < len(a); i++ {
		fmt.Println("i", i)
		fmt.Println("r", <-ch)
	}
}

func cast() {
	val, val2 := 15, 7
	resultFloat := divideIntToFloat(val, val2)
	fmt.Printf("divideIntToFloat %.1f\n", resultFloat)

	valString, val2String := "30", "15"
	resultInt := divideStringToInt(valString, val2String)
	fmt.Printf("divideStringToInt %d\n", resultInt)

	dataInput := os.Args[1]
	dataInput2 := os.Args[2]
	resultIntInput := divideStringToInt(dataInput, dataInput2)
	fmt.Printf("data input %d\n", resultIntInput)
}

func divideIntToFloat(val, val2 int) float32 {
	return float32(val / val2)
}

func divideStringToInt(val, val2 string) int {
	valInt, _ := strconv.Atoi(val)
	valInt2, _ := strconv.Atoi(val2)
	return valInt + valInt2
}
