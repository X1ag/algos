package main 

func moveZeroes(nums []int)  {
    if len(nums) == 0 {
     return 
    }

    slow := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[slow] = nums[slow], nums[i]

            slow++
        }
    }
}