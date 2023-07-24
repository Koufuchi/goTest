package printTest

import (
	"fmt"
	"sync"
)

func PrintByFor(start, end int) {
	if start > end {
		fmt.Println("start is great than end !")
	}

	for ; start <= end; start++ {
		fmt.Println(start)
	}
}

func PrintByRecu(start, end int) {
	if start > end {
		return
	}
	fmt.Println(start)
	PrintByRecu(start+1, end)
}

func PrintByRecuPoint(start, end int) {
	if start > end {
		fmt.Println("start is great than end !")
	}

	startPoint := &start
	endPoint := &end

	// 不能在一個函數內部聲明另一個函數，所以要先聲明後才能在匿名函數內呼叫
	var subFn func(*int, *int)
	subFn = func(start, end *int) {
		if *start > *end {
			return
		}

		fmt.Println(start, end) // 驗證是傳址非傳值
		fmt.Println(*start)
		*start++

		subFn(start, end)
	}

	subFn(startPoint, endPoint)
}

func PrintByChan(start, end int) {
	if start > end {
		fmt.Println("start is great than end !")
	}

	// chInt := make(chan int, 10) // 有緩衝代表不會每寫入一個內容就必須要先傳出去
	chInt := make(chan int)
	go printByChanSub(start, end, chInt)

	for num := range chInt {
		fmt.Println(num)
	}
}

func printByChanSub(start, end int, chInt chan int) {
	for ; start <= end; start++ {
		chInt <- start
		// fmt.Println("channel 長度: " + fmt.Sprint(len(chInt)))
	}

	close(chInt)
}

func PrintByChan2(start, end int) {

	chInt := make(chan int)

	printFn := func(chInt chan int) {
		for num := range chInt {
			fmt.Println(num)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	loopFn := func(start, end int, chInt chan int, wg *sync.WaitGroup) {
		for ; start <= end; start++ {
			chInt <- start
		}
		close(chInt)
		wg.Done()
	}

	go printFn(chInt)
	loopFn(start, end, chInt, &wg)
	wg.Wait()
}
