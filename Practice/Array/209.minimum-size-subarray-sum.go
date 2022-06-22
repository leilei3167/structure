package main

import (
	"math"
)

//https://leetcode.cn/problems/minimum-size-subarray-sum/
/*


给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0 。

。*/

//滑动窗口经典题,其实滑动窗口法本质也是双指针法的一种

/*
本题中实现滑动窗口，主要确定如下三点：
窗口内是什么？---之和>=target并且长度最小的连续子数组

如何移动窗口的起始位置？-----当窗口的值大于了target,就说明应该将起始位置移动以缩小窗口

如何移动窗口的结束位置？------结束位置就是数组的索引,即for循环的终止条件

解题的关键在于 窗口的起始位置如何移动,动窗口的精妙之处在于根据当前子序列和大小的情况，不断调节子序列的起始位置

*/

func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)          // 数组长度
	sum := 0                // 子数组之和
	result := math.MaxInt32 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况

	for j := 0; j < l; j++ { //移动右边界
		sum += nums[j]      //子数组之和
		for sum >= target { //子数组之和大于等于target,说明应该将起始位置移动以缩小窗口
			subLen := j - i + 1 //子数组长度
			if subLen < result {
				result = subLen
			}
			sum -= nums[i] //减去排出左边界的值
			i++
		}
	}
	if result == math.MaxInt32 { //设置一个极大的数,仅仅用于鉴别是否不存在符合条件的子数组
		return 0
	} else {
		return result
	}
}

func minSubArrayLen2(target int, nums []int) int {
	//经典滑动窗口
	//移动结束位置,当左边界和右边界元素和大于等于target时,说明应该移动左边界以缩小窗口
	left := 0
	result := len(nums) + 10000
	sum := 0
	for i := 0; i < len(nums); i++ {
		//元素和
		sum = sum + nums[i]
		for sum >= target {
			//左边界应该右移动缩小窗口,此时的子数组长度:
			subLen := i - left + 1
			if subLen < result {
				result = subLen
			}
			//排除上一个left的值
			sum = sum - nums[left]
			left++
		}
	}
	if result == len(nums)+10000 {
		return 0
	} else {
		return result
	}

}
