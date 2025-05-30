/*
Modify the below using "Share memory by communicating"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	dataCh := generateData(2, 100)
	primeCh := make(chan int)
	workerCount := 5

	for workerId := range workerCount {
		wg.Add(1)
		go primeWorker(workerId, wg, dataCh, primeCh)
	}

	// Using a channel (doneCh) for synchronization of the consumer (primeCh) goroutine
	doneCh := make(chan struct{})
	go func() {
		for primeNo := range primeCh {
			fmt.Println("Prime No :", primeNo)
		}
		close(doneCh)
	}()
	wg.Wait()
	close(primeCh)
	<-doneCh

	// Using another wait group (wg2) for synchronization of the consumer (primeCh) goroutine
	/*
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		go func() {
			for primeNo := range primeCh {
				fmt.Println("Prime No :", primeNo)
			}
			wg2.Done()
		}()
		wg.Wait()
		close(primeCh)
		wg2.Wait()
	*/
}

func primeWorker(workerId int, wg *sync.WaitGroup, dataCh <-chan int, primeCh chan<- int) {
	defer wg.Done()
	for no := range dataCh {
		fmt.Printf("primeWorker [id = %d] processing no = %d\n", workerId, no)
		if isPrime(no) {
			primeCh <- no
		}
	}
}

func generateData(start, end int) <-chan int {
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
