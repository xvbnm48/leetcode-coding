func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// Fill the array with values larger than amount
	// At most, denomination 1 coins would supply change for 'n' 'n' times
	mins := make([]int, amount+1)
	for i := range mins {
		mins[i] = amount + 1
	}

	// It takes 0 coins to make up the amount of 0
	mins[0] = 0

	for a := 1; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			// coin is larger than current amount
			if coins[c] > a {
				continue
			}

			// mins[a] is initially large enough to after being filled into slice above
			mins[a] = int(math.Min(float64(mins[a]), float64(mins[a-coins[c]]+1)))
		}
	}

	if mins[amount] > amount {
		return -1
	}

	return mins[amount]
}
