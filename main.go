package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/* This Go program generates random numbers and checks if they are prime numbers using concurrent goroutines.
It demonstrates the concepts of fan-out and fan-in patterns in Go concurrency.
*/

// take function takes n elements from a stream until done signal is received.
func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

// repeatFunc function repeats a function indefinitely until done signal is received.
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

// primeFinder function checks if a given number is prime or not.
func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

// fanIn function combines multiple channels into a single channel.
func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	// transfer function reads from each channel and forwards to fannedInStream.
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	// Start a goroutine for each channel.
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	// Close fannedInStream after all goroutines finish.
	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	// Define a function to generate random numbers.
	randNumFetcher := func() int {
		return rand.Intn(500000000)
	}
	randIntStream := repeatFunc(done, randNumFetcher)

	// fan out: Create multiple primeFinder goroutines.
	cpuCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, cpuCount)
	for i := 0; i < cpuCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in: Combine all primeFinder channels into a single channel.
	fannedInStream := fanIn(done, primeFinderChannels...)

	// Take first 10 prime numbers from the combined channel and print them.
	for rando := range take(done, fannedInStream, 10) {
		fmt.Println(rando)
	}

	fmt.Println(time.Since(start)) // Measure execution time
}
