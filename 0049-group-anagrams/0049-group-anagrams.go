// func groupAnagrams(strs []string) [][]string {
//     isAnagramGroup := make(map[string][]string)
    
//     for _, char := range strs {
//         isAnagramGroup[char]++
//     }
    
    
    
// }

func groupAnagrams(strs []string) [][]string {
    anagramGroups := make(map[string][]string)

    for _, str := range strs {
        // Buat representasi urutan karakter menggunakan penghitungan frekuensi karakter
        charCount := make([]int, 26)
        for _, char := range str {
            charCount[char-'a']++
        }

        // Konversi representasi urutan karakter menjadi string
        sortedStr := ""
        for i := 0; i < 26; i++ {
            if charCount[i] > 0 {
                sortedStr += string('a' + i) + strconv.Itoa(charCount[i])
            }
        }

        // Gunakan representasi urutan karakter sebagai kunci map
        // dan tambahkan string asli ke dalam grup yang sesuai
        anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
    }

    // Konversi map ke dalam slice hasil akhir
    var result [][]string
    for _, group := range anagramGroups {
        result = append(result, group)
    }

    return result
}
