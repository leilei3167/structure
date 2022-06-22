package main

//https://leetcode.cn/problems/binary-search/

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

*/

/*
 题解:
 这道题目的前提是数组为有序数组，同时题目还强调数组中无重复元素，因为一旦有重复元素，使用二分查找法返回的元素下标可能不是唯一的，
 这些都是使用二分法的前提条件，当大家看到题目描述满足如上条件的时候，可要想一想是不是可以用二分法了
*/

/* 关键是一开始就定义好开闭,并且全部统一 */

//方法一:左闭右闭,[left,target],每一个元素都是可达的下标
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left //中间值是左边值加上左边值和右边值差的一半
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

//方法二,左闭右开,mid作为右边界,起始值为len(nums),mid作为不可达的下标
func search2(nums []int, target int) int {
	left, right := 0, len(nums) //right是长度,不可达
	for left < right {          //不能等于,因为右边界不可能相等
		mid := (right-left)/2 + left //中间
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { //说明下标一定在mid左边
			right = mid
		} else {
			left = mid + 1
		}

	}
	return -1
}
