package main

import "fmt"

// 求解一个有序数组，过滤掉重复的数字，返回过滤后的长度
// [1, 2, 3, 3, 3, 4, 4, 5, 7, 8]
// => [1, 2, 3, 4, 5, 6, 7, 8]

// 思路：需要有一个索引下标，标识存储不重复数字的最新下标，然后往后依次跟这个下标的前一个数字对比，如果不相同，则不是重复的数字，将下标对应的数换为最新的，下标往前移动，依次重复。

func main() {
	sli := []int{1, 2, 2, 3, 3, 3, 4, 6, 6, 7, 10}
	fmt.Println("init => ", sli)
	length := removeDuplicates(sli)
	fmt.Println("length => ", length)
	fmt.Println("removed => ", sli[:length])
}

func removeDuplicates(nums []int) (length int) {
	if len(nums) == 0 {
		return 0
	}
	length = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length-1] {
			nums[length] = nums[i]
			length++
		}
	}
	return
}
