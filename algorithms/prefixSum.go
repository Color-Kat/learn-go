package main

import "fmt"

func prefixSum(nums []int) {
	prefixes := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixes[i+1] = prefixes[i] + nums[i]
	}

	fmt.Println(prefixes)
}

func main() {
	prefixSum([]int{1, 2, 3, 4, 5})

}
