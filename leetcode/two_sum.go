/*
 * @Date: 2021-08-13 07:24:26
 * @LastEditors: seven sun
 * @LastEditTime: 2021-08-13 14:29:32
 * @FilePath: /面试题/leetcode/two_sum.go
 */
package leetcode

func Sum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}
