func containsNearbyDuplicate(nums []int, k int) bool {
    numMap := make(map[int]int)
    
    for i, num := range nums {
        if idx, found := numMap[num]; found {
            if i-idx <= k {
                return true
            }
        }
         numMap[num] = i
    }
   return false
}