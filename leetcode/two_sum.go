/*
 * @Date: 2021-08-13 07:24:26
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-08-13 13:56:09
 * @FilePath: /interview/leetcode/two_sum.go
 */
package leetcode

func Sum(nums []int, target int) []int{
   for (i:=0;i++;i<len(nums)){
	   for (j:=0;j++;j<len(nums)){
		   if nums[i]+nums[j]=target{
		   return []int{nums[i],nums[j]}
		   }
   }
  }
return []int{}
}

