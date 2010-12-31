// demonstartion of channels in go
package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'.
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {

	// make function is used to create a new channel
	ch := make(chan int)  // Create a new channel.
	
	// Start generate() as a subprocess.
	// 'go' statement starts generate(ch) function in a
	// independent thread
	
	// this is what is called a goroutine
	
	go generate(ch)
	
	     
	//for j := 1; j<20; j++ {
		// initialize prime with what comes from ch
		prime := <-ch
		fmt.Print(prime, "\n")
		
		// next value of channel goes into another variable
		prime1 := <-ch
		fmt.Print(prime1, "\n")
		
	//}
}

func main() {
	sieve()
}
