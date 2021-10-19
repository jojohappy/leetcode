/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	index := -1

	for l <= r {
		mid := l + (r-l)/2
		v := nums[mid]
		if v == target {
			return mid
		}

		if v > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}

// @lc code=end

