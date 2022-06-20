package main

//https://leetcode.cn/problems/squares-of-a-sorted-array/

//有序数的平方
//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

//最暴力的方式是每个元素平方后排序

//记住数组是有序的,只不过负数的最小值平方后就可能是最大值了,因此进行最左和最右的值平方后进行比较,较大的放入到新数组从右向左放
//双指针法,一个指针指向开始,一个指针指向末尾
//时间复杂度O(n)
func sortedSquares(nums []int) []int {
	n := len(nums)
	//i,j分别代表nums的最左边和最右边
	i, j, k := 0, n-1, n-1
	//创建一个新数组
	res := make([]int, n)
	for i <= j { //k>=0
		//比较平方的大小
		si, sj := nums[i]*nums[i], nums[j]*nums[j]
		if si > sj {
			res[k] = si //左边更大则赋值给新数组最右侧,并且i++右移动
			i++
		} else {
			//右边更大,赋值后左移
			res[k] = sj
			j--
		}
		k--
	}
	return res

}
