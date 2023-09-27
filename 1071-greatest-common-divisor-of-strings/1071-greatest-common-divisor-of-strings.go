func gcdOfStrings(str1 string, str2 string) string {
    if str1+str2 != str2+str1 {
        return ""
    }
    
    // Mencari faktor pembagi terbesar dari panjang kedua string
    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }(len(str1), len(str2))
    
    return str1[:gcd]
}