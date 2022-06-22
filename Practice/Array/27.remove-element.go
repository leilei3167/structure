package main

import "fmt"

//https://leetcode.cn/problems/remove-element/

//双for循环可暴力解,但是快慢指针是更优雅的方式
func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ { //遍历数组,即是快指针移动,当出现val时慢指针忽略,将不等于val的元素放入慢指针数组
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}

	}
	fmt.Printf("%#v\n", nums)
	return n
}
