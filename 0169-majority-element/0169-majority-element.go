func majorityElement(nums []int) int {
    hitung := 0 
    kandidat := 0
    
    for _, num := range nums {
        if hitung == 0 {
            kandidat = num
        } 
        
        if kandidat == num {
            hitung++
        } else {
            hitung--
        }
    }
    return kandidat
}

// Dalam solusi ini, kita menggunakan pendekatan "Boyer-Moore Voting Algorithm". Ide utamanya adalah menghitung suara positif dan suara negatif untuk elemen-elemen dalam array. Kita mulai dengan menganggap elemen pertama sebagai kandidat mayoritas, dan setiap kali kita menemukan elemen yang sama, kita meningkatkan suara positif. Jika kita menemukan elemen yang berbeda, kita mengurangi suara positif. Jika suara positif menjadi 0, kita mengganti kandidat dengan elemen saat ini.

// Karena elemen mayoritas muncul lebih dari ⌊n / 2⌋ kali, dan elemen lain tidak dapat mengimbangi suara mayoritas, maka kandidat mayoritas akan tetap ada setelah semua elemen dihitung. Oleh karena itu, setelah selesai perulangan, kandidat mayoritas adalah elemen mayoritas yang kita cari.