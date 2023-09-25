func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sToT := make(map[byte]byte)
    tToS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        charS := s[i]
        charT := t[i]

        // Periksa pemetaan dari s ke t
        if mappedT, found := sToT[charS]; found {
            if mappedT != charT {
                return false
            }
        } else {
            sToT[charS] = charT
        }

        // Periksa pemetaan dari t ke s
        if mappedS, found := tToS[charT]; found {
            if mappedS != charS {
                return false
            }
        } else {
            tToS[charT] = charS
        }
    }

    return true
}
