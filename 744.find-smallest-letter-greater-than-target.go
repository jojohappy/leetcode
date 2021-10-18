/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */
package main

import "fmt"

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	length := len(letters)
	l := 0
	r := length - 1
	for {
		if l > r {
			break
		}
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return letters[l%length]
}

// @lc code=end

func main() {
	letters := "cfjg"
	target := nextGreatestLetter([]byte(letters), 'c')
	fmt.Printf("%c\n", target)
}
