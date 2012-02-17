package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
	"time"
)

func main() {
	high := uint64(1000000000)
	sieve := primegen.New()
	start := time.Now()
	count := sieve.Count(high)
	end := time.Now()
	fmt.Printf("%d primes up to %d\n", count, high)
	fmt.Printf("Overall seconds: approximately %.5f\n", float64(end.Sub(start))/1e9)
}
