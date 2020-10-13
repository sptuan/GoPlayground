package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 5)
	fmt.Printf("[Chan] Create a chan int\n")
	//chE := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(4)

	// basic go chan usage test
	go func() {
		defer wg.Done()
		fmt.Printf("[Send] A send go func start running.")
		for i := 0; i < 100; i++ {
			ch1 <- i
			fmt.Printf("[Send] %x\n", i)
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("[Recv] A recv go func start running.")
		for {
			e, ok := <-ch1
			if ok == false {
				break
			}
			fmt.Printf("[Recv] %x\n", e)
		}
		fmt.Printf("[Recv] Channel Closed.\n")
		//close(chE)
	}()

	// other usage test
	ch2 := make(chan int, 5)
	fmt.Printf("[Chan2] Create a chan int 2\n")
	go func() {
		defer wg.Done()
		fmt.Printf("[Send2] A send go func start running.")
		for i := 0; i < 100; i++ {
			ch2 <- i
			fmt.Printf("[Send2] %x\n", i)
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("[Recv2] A recv go func start running.")
		for e := range ch2 {
			fmt.Printf("[Recv2] recv a e: %d\n", e)
		}
		fmt.Printf("[Recv2] Channel Closed.\n")
	}()

	fmt.Println("[Task] All go routine created.")
	//_, ok := <- chE
	//if ok==false {
	//	os.Exit(0)
	//}
	wg.Wait()
	fmt.Println("[Main] All wait group finished.")
}
