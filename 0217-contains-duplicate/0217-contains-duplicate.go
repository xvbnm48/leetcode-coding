// func containsDuplicate(nums []int) bool {
//   sort.Ints(nums)
//     for i := 1; i < len(nums); i++ {
//         if nums[i] == nums[i-1] {
//             return true
//         }
//     }
//     return false
// }

func containsDuplicate(nums []int) bool {
    // make seen for store value in slice of array bool
    seen := make(map[int]bool)
    
    // for loop for range of nums 
    for _, num := range nums {
        if seen[num] {
            return true 
        } 
        seen[num] = true
    }
    return false
}