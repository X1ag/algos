package main

func maxArea(height []int) int {
	if len(height) == 0 {
			return 0
	}

	l := 0
	r := len(height) - 1
	maxArea := 0
	currLineArea := 0

	for i := 0; i < len(height); i++ {
			if height[l] <= height[r] {
					if l > r {
							currLineArea = l - r
					} else {
							currLineArea = r - l
					}
					currArea := height[l] * currLineArea
					if maxArea < currArea {
							maxArea = currArea
					}
					l++
			} else {
					if l > r {
							currLineArea = l - r
					} else {
							currLineArea = r - l
					}
					currArea := height[r] * currLineArea
					if maxArea < currArea {
							maxArea = currArea
					}
					r--
			}
	} 
	return maxArea
}