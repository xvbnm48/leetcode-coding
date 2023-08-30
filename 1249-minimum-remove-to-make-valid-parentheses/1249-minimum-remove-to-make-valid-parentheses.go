func minRemoveToMakeValid(s string) string {
    stack := []int{}
    chars := []byte(s)
    
    // Tandai indeks karakter yang perlu dihapus
    for i, c := range chars {
        if c == '(' {
            stack = append(stack, i)
        } else if c == ')' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1] // Pop dari stack
            } else {
                chars[i] = '*' // Tandai karakter yang perlu dihapus
            }
        }
    }
    
    // Tandai karakter dalam stack sebagai perlu dihapus
    for _, idx := range stack {
        chars[idx] = '*'
    }
    
    // Buat string hasil tanpa karakter yang ditandai untuk dihapus
    result := []byte{}
    for _, c := range chars {
        if c != '*' {
            result = append(result, c)
        }
    }
    
    return string(result)
}
