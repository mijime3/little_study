package main

import "fmt"

func is_prime(n int) bool{
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false;
		}
	}
	return true;
}

func main() {
	for i := 2; i <= 50000; i++ {
		if is_prime(i) {
			fmt.Printf("[%d]\n", i)
		}
	}
}	
