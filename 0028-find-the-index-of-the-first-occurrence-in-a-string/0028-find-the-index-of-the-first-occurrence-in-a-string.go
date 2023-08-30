func strStr(haystack string, needle string) int {
    hLen := len(haystack)
    nLen := len(needle)
    
    if nLen == 0 {
        return 0
    }
    
    for i := 0; i <= hLen - nLen; i++{
        if haystack[i:i+nLen] == needle {
            return i
        }
    }
    
    return -1
}
