package main 

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    l := 0 
    for i := 1; i < len(nums); i++ {
        if nums[l] != nums[i] {
            l++
            nums[l] = nums[i]
        }   
    }
    return l + 1
}