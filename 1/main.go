package main

import "fmt"

func main() {
	var res int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}

	fmt.Println(res)
}
