func hitungKuadratDigit(n int) int {
    total := 0
    for n > 0 {
        digit := n % 10
        total += digit * digit
        n /= 10
    }
    return total
}

func isHappy(n int) bool {
    lambat, cepat := n, n

    for {
        lambat = hitungKuadratDigit(lambat)
        cepat = hitungKuadratDigit(hitungKuadratDigit(cepat))

        if lambat == 1 {
            return true
        }

        if lambat == cepat {
            return false
        }
    }
}
