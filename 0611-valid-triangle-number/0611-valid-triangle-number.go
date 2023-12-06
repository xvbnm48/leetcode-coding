func triangleNumber(nums []int) int {
    n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)

	count := 0

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				// Jika dapat membentuk segitiga, maka dapat membentuk segitiga dengan semua bilangan di antara left dan right.
				count += right - left
				right--
			} else {
				// Jika tidak dapat membentuk segitiga, maka coba dengan meningkatkan left.
				left++
			}
		}
	}

	return count
}