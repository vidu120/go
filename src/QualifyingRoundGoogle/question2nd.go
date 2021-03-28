package main

//this function calculates the minimum cost for the mural operation which is described in the question
func calculateMinimalCost(mural string, X, Y int) int {
	if len(mural) == 0 || len(mural) == 1 {
		return 0
	}
	cost := 0
	var prevCharacter = ' '
	var postCharacter = ' '
	for i := 0; i < len(mural); i++ {
		if mural[i] == '?' {
			if i-1 >= 0 {
				prevCharacter = rune(mural[i-1])
			}
			for i < len(mural) && mural[i] == '?' {
				i++
			}
			if i != len(mural) {
				postCharacter = rune(mural[i])
			}
			if prevCharacter == 'C' && postCharacter == 'J' {
				cost += X
			} else if prevCharacter == 'J' && postCharacter == 'C' {
				cost += Y
			}
			prevCharacter = ' '
			postCharacter = ' '
			i--
		}
	}

	//calculate the original cost
	for i := 0; i < len(mural)-1; i++ {
		if mural[i:i+2] == "CJ" {
			cost += X
		} else if mural[i:i+2] == "JC" {
			cost += Y
		}
	}
	return cost
}
