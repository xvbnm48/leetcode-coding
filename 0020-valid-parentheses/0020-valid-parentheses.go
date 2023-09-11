// func isValid(s string) bool {
//     l := list.New()

//     for _, v := range s {
//        if v == '{' || v == '(' || v == '[' {
// 			l.PushFront(v)
// 		} else if l.Len() > 0 {
// 			e := l.Front()
// 			if v == '}' && e.Value != '{' {
// 				return false
// 			}
// 			if v == ')' && e.Value != '(' {
// 				return false
// 			}
// 			if v == ']' && e.Value != '[' {
// 				return false
// 			}
// 			l.Remove(e)
// 		} else {
// 			return false
// 		}
//     }
//     if l.Len() >0 {
//         return false
//     }
//     return true
// }


func isValid(s string) bool {
    stack := []rune{} // Membuat tumpukan kosong menggunakan tipe rune

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            // Jika karakter buka kurung, masukkan ke dalam tumpukan
            stack = append(stack, char)
        } else {
            // Jika karakter tutup kurung, periksa apakah tumpukan kosong
            if len(stack) == 0 {
                return false
            }
            
            // Ambil elemen terakhir dari tumpukan
            top := stack[len(stack)-1]
            
            // Periksa apakah karakter tutup kurung sesuai dengan yang diharapkan
            if (char == ')' && top == '(') || (char == '}' && top == '{') || (char == ']' && top == '[') {
                // Pop elemen terakhir dari tumpukan
                stack = stack[:len(stack)-1]
            } else {
                return false // Tidak sesuai, string tidak valid
            }
        }
    }

    // String dianggap valid jika tumpukan kosong pada akhirnya
    return len(stack) == 0
}