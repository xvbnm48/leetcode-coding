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
    seen := make(map[int]bool)
    
    for _, num := range nums {
        if seen[num] {
            return true 
        } 
        seen[num] = true
    }
    return false
}