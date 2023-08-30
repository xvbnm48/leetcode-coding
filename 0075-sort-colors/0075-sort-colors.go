func sortColors(nums []int)  {
    hitung := [3]int{}
    
    // isi var hitung dengan semua angka di parameter nums
    for _, num := range nums {
        hitung[num]++
    }
    
    awal := 0
    
    for warna := 0; warna <= 2; warna++ {
        for hitung[warna] > 0 {
            nums[awal] = warna 
            awal++
            hitung[warna]--
        } 
    }
}