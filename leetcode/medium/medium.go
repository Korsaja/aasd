package medium

// Runtime: 9 ms, faster than 90.91%
// Memory Usage: 7.6 MB, less than 16.88%
func minPartitions(s string) int {
	var max int
	for _, r := range s {
		max = getMax(max, runeToInt(r))
	}
	return max
}

// Runtime: 35 ms, faster than 45.54%
// Memory Usage: 5.1 MB, less than 79.46%
func minOperations(boxes string) []int {
	l := len(boxes)
	result := make([]int, l)
	for i, box := range boxes {
		if runeToInt(box) == 1 {
			for j := l - 1; j >= 0; j-- {
				result[j] += abs(i - j)
			}
		}
	}
	return result
}

// Runtime: 35 ms, faster than 28.00%
// Memory Usage: 3.8 MB, less than 40.00%
func executeInstructions(n int, startPos []int, operations string) []int {
	var step = func(op uint8, x, y int) (bool, int, int) {
		isOver := func(x, y int) bool {
			return x < 0 || y < 0 || x > n-1 || y > n-1
		}
		switch op {
		case 'R':
			x++
		case 'L':
			x--
		case 'D':
			y++
		case 'U':
			y--
		}
		return isOver(x, y), x, y
	}

	l := len(operations)
	var res = make([]int, l)
	for i := range operations {
		var isOver bool
		x, y := startPos[1], startPos[0]
		count := 0
		for j := i; j < l; j++ {
			isOver, x, y = step(operations[j], x, y)
			if isOver {
				break
			}
			count++
			if j == l-1 {
				break
			}
		}
		res[i] = count
	}
	return res
}

// helpers
func runeToInt(r rune) int {
	return int(r - '0')
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
