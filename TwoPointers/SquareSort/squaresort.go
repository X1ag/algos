package main

import (
	"fmt"
	"sort"
)

func sortedSquaresStupid(nums []int) []int {
		for i := 0; i < len(nums); i++ {
				nums[i] = nums[i] * nums[i]
		}
		sort.Slice(nums, func(i, j int) bool {
    	return nums[i] < nums[j]
	})
	
	return nums 
}

func testCases(nums []int, copyNums []int) bool {
		sortedNums := sortedSquaresStupid(copyNums)
		for i := 0; i < len(sortedNums); i++ {
			if nums[i] != sortedNums[i]{
				fmt.Println("not sorted", nums)
				panic("not sorted")
			}
		}
		return true 
}

func sortedSquaresSmart(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }

    result := make([]int, len(nums))
    l := 0
    r := len(nums) - 1

    for i := len(nums) - 1; i >= 0; i-- {
        leftValue := nums[l] * nums[l]
        rightValue := nums[r] * nums[r]
        if leftValue <= rightValue {
            result[i] = rightValue
            r--
        } else {
            result[i] = leftValue
            l++
        }
    }

	if testCases(result, nums) {
    	fmt.Println("sorted:", result)
  	  return result 
	}
		return []int{}
}


func main() {
	sortedSquaresSmart([]int{-4,-1,0,3,10})
}