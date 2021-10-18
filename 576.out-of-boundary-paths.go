/*
 * @lc app=leetcode id=576 lang=golang
 *
 * [576] Out of Boundary Paths
 */
package main

import (
	"fmt"
)

// @lc code=start
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	paths := 0
	grids := make([][]int, m)
	for idx, _ := range grids {
		grids[idx] = make([]int, n)
	}

	for idx, rows := range grids {
		for cellIdx, cell := range rows {
			if idx == 0 || idx == (m-1) {
				paths += cell
			} else if cellIdx == 0 || cellIdx == (n-1) {
				paths += cell
			}
		}
	}
	return paths
}

// @lc code=end

func main() {
	m := 1
	n := 3
	maxMove := 3
	startRow := 0
	startColumn := 1
	output := findPaths(m, n, maxMove, startRow, startColumn)
	fmt.Println(output)
}
