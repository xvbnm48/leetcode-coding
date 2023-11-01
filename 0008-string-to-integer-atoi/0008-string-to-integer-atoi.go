import "math"

func myAtoi(s string) int {
    result := 0
    sign := 1
    i := 0

    // Langkah 1: Lewati spasi awal
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // Langkah 2: Cek tanda positif atau negatif
    if i < len(s) && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    // Langkah 3: Baca digit-digit berikutnya hingga karakter yang bukan digit
    for i < len(s) && isDigit(s[i]) {
        digit := int(s[i] - '0')

        // Langkah 4: Perbarui hasil dengan digit yang dibaca
        if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
            if sign == 1 {
                return math.MaxInt32
            } else {
                return math.MinInt32
            }
        }
        result = result*10 + digit
        i++
    }

    // Langkah 5: Clamp hasil ke dalam rentang [-2^31, 2^31-1]
    result *= sign

    if result > math.MaxInt32 {
        return math.MaxInt32
    } else if result < math.MinInt32 {
        return math.MinInt32
    }

    return result
}

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}
