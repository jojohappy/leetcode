/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

package main

import (
	"math"
)

// @lc code=start
type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	lastIndex := len(this.stack) - 1
	val := this.stack[lastIndex]
	this.stack = this.stack[:lastIndex]
	if val == this.min {
		this.resetMinVal()
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) resetMinVal() {
	min := math.MaxInt32
	for _, v := range this.stack {
		if min > v {
			min = v
		}
	}
	this.min = min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
