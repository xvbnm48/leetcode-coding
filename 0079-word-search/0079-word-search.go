func exist(board [][]byte, word string) bool {
    // Mendapatkan panjang dan lebar dari matriks
    m, n := len(board), len(board[0])
    
    // Fungsi rekursif untuk melakukan pencarian
    var search func(i, j, k int) bool
    search = func(i, j, k int) bool {
        // Jika kita sudah mencapai akhir kata, kembalikan true
        if k == len(word) {
            return true
        }
        
        // Jika posisi saat ini di luar batas matriks,
        // atau huruf di sel tidak sesuai dengan huruf yang sedang dicari, kembalikan false
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
            return false
        }
        
        // Simpan karakter yang ada di sel saat ini
        temp := board[i][j]
        
        // Tandai sel saat ini sebagai karakter yang tidak valid (agar tidak digunakan lagi)
        board[i][j] = '#'
        
        // Coba bergerak ke sel-sel tetangga dalam empat arah (atas, bawah, kiri, kanan)
        found := search(i+1, j, k+1) || search(i-1, j, k+1) || search(i, j+1, k+1) || search(i, j-1, k+1)
        
        // Kembalikan karakter sel ke nilai aslinya
        board[i][j] = temp
        
        return found
    }

    // Cari huruf pertama dari kata pada setiap sel dalam matriks
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] && search(i, j, 0) {
                return true
            }
        }
    }
    
    return false
}
