package main

import "fmt"

// 数字能重复少于2次，大于2次的需要去掉重复的数字
// [1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 4, 5, 6, 7]
// => [1, 2, 2, 3, 3, 4, 4, 5, 6, 7]

// 思路：
// 把一次换成出现两次，同样可以实现
//
// 扩展：
// 理论上可以将1次换为n次，变成通用

func main() {
	sli := []int{1, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 6, 7, 7, 7, 7}
	fmt.Println(sli)

	length := removedDuplicatedTwice(sli)
	fmt.Println("length => ", length)
	fmt.Println("removed => ", sli[:length])
}

// 可以重复多少次
func removedDuplicatedTwice(nums []int) int {
	return removedDuplicatedOccur(nums, 2)
}

// 通用版本
func removedDuplicatedOccur(nums []int, occur int) (length int) {
	if len(nums) <= occur {
		return len(nums)
	}

	length = occur
	for i := occur; i < len(nums); i++ {
		if nums[i] != nums[length-occur] {
			nums[length] = nums[i]
			length++
		}
	}
	return
}
