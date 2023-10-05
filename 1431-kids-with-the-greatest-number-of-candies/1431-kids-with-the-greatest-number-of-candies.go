func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := 0
    for _ , permen := range candies {
        if permen > maxCandies {
            maxCandies = permen
        }
    }
    
    result := make([]bool, len(candies))
    
    for i, permen := range candies {
        result[i] = permen+extraCandies >= maxCandies
    }
    return result
}   