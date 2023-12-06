func canPlaceFlowers(flowerbed []int, n int) bool {
     // Variabel count yang menunjukkan jumlah bunga yang sudah ditanam
    count := 0
    // Iterasi pada array flowerbed dengan indeks i
    for i := 0; i < len(flowerbed); i++ {
        // Jika flowerbed[i] sama dengan 0, yaitu plot kosong
        if flowerbed[i] == 0 {
            // Jika i sama dengan 0 atau flowerbed[i-1] sama dengan 0, yaitu plot sebelumnya kosong atau tidak ada
            if i == 0 || flowerbed[i-1] == 0 {
                // Jika i sama dengan panjang array dikurangi 1 atau flowerbed[i+1] sama dengan 0, yaitu plot sesudahnya kosong atau tidak ada
                if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
                    // Ubah flowerbed[i] menjadi 1, yaitu menanam bunga
                    flowerbed[i] = 1
                    // Tambah count dengan 1, yaitu menambah jumlah bunga yang sudah ditanam
                    count++
                }
            }
        }
    }
    // Bandingkan count dengan n, dan kembalikan hasilnya
    return count >= n
}