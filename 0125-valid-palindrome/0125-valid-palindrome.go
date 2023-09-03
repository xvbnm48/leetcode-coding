func isPalindrome(s string) bool {
    cleanStr  := make([]rune, 0, len(s))
    
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9'){
            cleanStr  = append(cleanStr , char)
        } else if char >= 'A' && char <= 'Z' {
            cleanStr  = append(cleanStr , char+'a'-'A')
        }
    }
    
    for i := 0; i < len(cleanStr)/2; i++{
        if cleanStr[i] != cleanStr[len(cleanStr )-i-1] {
            return false
        }
    }
    return true  
}