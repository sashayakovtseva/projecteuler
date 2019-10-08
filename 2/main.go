package main

import "fmt"

func main() {
	var res int64
	prev := int64(1)
	curr := int64(1)

	for next := prev + curr; next < 4000000; next = prev + curr {
		if next&1 == 0 {
			res += next
		}
		prev = curr
		curr = next
	}

	fmt.Println(res)
}
