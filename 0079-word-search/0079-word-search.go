func exist(board [][]byte, word string) bool {
    // Membuat array bool untuk melacak sel-sel yang sudah dikunjungi
    visited := make([][]bool, len(board))
    for i := range visited {
        visited[i] = make([]bool, len(board[0]))
    }

    // Fungsi rekursif untuk melakukan pencarian
    var search func(i, j, k int) bool
    search = func(i, j, k int) bool {
        // Jika kita sudah mencapai akhir kata, kembalikan true
        if k == len(word) {
            return true
        }
        
        // Jika posisi saat ini di luar batas matriks atau sel telah dikunjungi,
        // atau huruf di sel tidak sesuai dengan huruf yang sedang dicari, kembalikan false
        if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[k] {
            return false
        }
        
        // Tandai sel saat ini sebagai sudah dikunjungi
        visited[i][j] = true
        
        // Coba bergerak ke sel-sel tetangga dalam empat arah (atas, bawah, kiri, kanan)
        found := search(i+1, j, k+1) || search(i-1, j, k+1) || search(i, j+1, k+1) || search(i, j-1, k+1)
        
        // Setelah selesai pencarian dalam satu arah, kembalikan status visited ke false
        visited[i][j] = false
        
        return found
    }

    // Cari huruf pertama dari kata pada setiap sel dalam matriks
    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] && search(i, j, 0) {
                return true
            }
        }
    }
    
    return false
}
