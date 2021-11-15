package logic

func SolveSudoku(input *[9][9]int) bool {
	if !hasEmptyCell(input) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if input[i][j] == 0 {
				for num := 9; num >= 1; num-- {
					input[i][j] = num
					if isBoardValid(input) {
						if SolveSudoku(input) {
							return true
						}
						input[i][j] = 0
					} else {
						input[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(input *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if input[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(input *[9][9]int) bool {

	//check duplicates number in row
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[input[i][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates number in column
	for j := 0; j < 9; j++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[input[col][j]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates number in square 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[input[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
