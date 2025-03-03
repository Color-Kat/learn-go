package main

import (
	"fmt"
	"strings"
)

func lastWordLength(str string) int {
	length := 0
	for i := len(str) - 1; i > 0; i-- {
		if length == 0 && str[i] == ' ' {
			continue
		}

		if str[i] != ' ' {
			length++
		} else {
			break
		}
	}

	return length
}

func increaseSliceNumber(slice []int) []int {
	// Add new 0 digit at slice start
	slice = append([]int{0}, slice...)

	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < 9 {
			slice[i]++
			break
		} else {
			slice[i] = 0
		}
	}

	// No first digit
	if slice[0] == 0 {
		return slice[1:]
	}

	return slice
}

func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1

	// Search among all rows
	for top <= bottom {
		middleRow := (top + bottom) / 2

		if target < matrix[middleRow][0] {
			bottom = middleRow - 1
			continue
		}

		if target > matrix[middleRow][len(matrix[middleRow])-1] {
			top = middleRow + 1
			continue
		}

		// Search among one row
		left, right := 0, len(matrix[middleRow])-1

		for left <= right {
			middleCol := (left + right) / 2

			if matrix[middleRow][middleCol] == target {
				return true
			} else if matrix[middleRow][middleCol] < target {
				left = middleRow + 1
			} else {
				right = middleRow - 1
			}
		}

		break
	}

	return false
}

func lengthOfLongestSubstring(s string) int {
	l, r, maxStreak := 0, 0, 0

	for _, char := range s {
		repeatIndex := strings.LastIndex(s[l:r], string(char))

		if repeatIndex >= 0 {
			l = l + repeatIndex + 1
		}

		r++

		if maxStreak < (r - l) {
			maxStreak = r - l
		}
	}

	return maxStreak
}

func main() {
	// fmt.Println(lastWordLength("I love typescript                 "))

	// fmt.Println(increaseSliceNumber([]int{9, 9}))
	// fmt.Println(increaseSliceNumber([]int{9}))
	// fmt.Println(increaseSliceNumber([]int{9, 9, 9, 9}))

	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))

	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring("aa"))
	// fmt.Println(lengthOfLongestSubstring("aaca"))
	// fmt.Println(lengthOfLongestSubstring(""))
	// fmt.Println(lengthOfLongestSubstring(" "))
	// fmt.Println(lengthOfLongestSubstring("aab"))
	// fmt.Println(lengthOfLongestSubstring("dvdf"))
	// fmt.Println(lengthOfLongestSubstring("anviaj"))
	// fmt.Println(lengthOfLongestSubstring("ohomm"))
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// [1,2,3,4,9]
