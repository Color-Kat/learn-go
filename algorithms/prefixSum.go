package main

import "fmt"

func prefixSum(nums []int, from, to int) int {
	prefixes := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixes[i+1] = prefixes[i] + nums[i]
	}

	return prefixes[to+1] - prefixes[from]
}

func main() {
	fmt.Println(prefixSum([]int{1, 2, 3, 4, 5}, 2, 4))

}
