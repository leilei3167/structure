package main

//https://leetcode.cn/problems/spiral-matrix-ii/

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
1<=n<=20
*/

/*
本题是经典面试题,填充的过程为:
填充上行从左到右
填充右列从上到下
填充下行从右到左
填充左列从下到上

非常容易循环边界条件出错,一定要坚持循环条件的一致性(如全部左闭右开),以n=3为例,每一条边的最后一个元素都应属于下一个方向
的起始,因此适用于左闭右开的循环条件
*/
func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1 //每条边的索引
	left, right := 0, n-1

	num := 1 //要填充的数字,从1开始到n的平方
	tar := n * n

	matrix := make([][]int, n) //每条边长为n
	//1.初始化二维数组
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	//从左上角开始
	for num <= tar { //直到达到n*n个元素填充完毕
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		//第二层,竖向
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++

	}
	return matrix

}
