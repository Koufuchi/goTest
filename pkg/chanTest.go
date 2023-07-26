package pkg

import (
	"fmt"
	"reflect"
	"time"
)

func FetchNumber(num int) chan int {
	r := make(chan int)
	go func() {
		time.Sleep(time.Duration(num) * time.Second)
		fmt.Println(num)
		r <- num * num
	}()
	return r
}

func Start() {
	fmt.Println("Fetching numbers ...")
	val1 := FetchNumber(1)

	if reflect.TypeOf(val1).Kind() == reflect.Int {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

	fmt.Println(2)
	fmt.Println(<-FetchNumber(3))
}
