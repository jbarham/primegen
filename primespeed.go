package main

import (
	"fmt"
	"time"
	"github.com/jbarham/primegen.go"
)

func main() {
	high := uint64(1000000000)
	sieve := primegen.New()
	start := time.Nanoseconds()
	count := sieve.Count(high)
	end := time.Nanoseconds()
	fmt.Printf("%d primes up to %d\n", count, high)
	fmt.Printf("Overall seconds: approximately %.5f\n", float64(end-start)/1e9)
}
