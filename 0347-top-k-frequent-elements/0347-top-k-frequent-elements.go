func topKFrequent(nums []int, k int) []int {
    // Hitung frekuensi setiap elemen
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// Temukan k elemen dengan frekuensi tertinggi
	result := make([]int, 0, k)
	for k > 0 {
		maxFreq := 0
		maxNum := 0

		// Cari elemen dengan frekuensi tertinggi
		for num, freq := range frequency {
			if freq > maxFreq {
				maxFreq = freq
				maxNum = num
			}
		}

		// Tambahkan elemen tersebut ke dalam hasil
		result = append(result, maxNum)

		// Hapus elemen tersebut dari map frekuensi
		delete(frequency, maxNum)

		k--
	}

	return result
}