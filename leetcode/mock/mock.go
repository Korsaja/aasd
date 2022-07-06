package mock

import (
	"math"
)

func diagonalSum(mat [][]int) int {
	sum := 0
	mid := 0

	for i, j := 0, len(mat)-1; i < len(mat) && j >= 0; i, j = i+1, j-1 {
		if i == j {
			mid = mat[i][j]
			continue
		}
		sum += mat[i][i]
		sum += mat[i][j]
	}

	return sum + mid
}
func dieSimulator(n int, rollMax []int) int {
	mod := int(math.Pow(10, 9)) + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 7)
	}
	for i := range dp[1] {
		dp[1][i] = 1
	}

	dp[1][6] = 6
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j <= 5; j++ {
			dp[i][j] = dp[i-1][6]
			if i == rollMax[j]+1 {
				dp[i][j] -= 1
			} else if i > rollMax[j]+1 {
				dp[i][j] -= dp[i-rollMax[j]-1][6] - dp[i-rollMax[j]-1][j]
				dp[i][j] = (dp[i][j] + mod) % mod
			}
			sum += dp[i][j]
			sum %= mod
		}
		dp[i][6] = sum
	}
	return dp[n][6]
}

func canFormArray(arr []int, pieces [][]int) bool {
	search := func(n int) int {
		for i, piece := range pieces {
			if piece[0] == n {
				return i
			}
		}
		return -1
	}

	var result []int
	for _, value := range arr {
		i := search(value)
		if i != -1 {
			result = append(result, pieces[i]...)
		}
	}
	return equalIntSlice(arr, result)
}

func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}
	var counter = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	var odd = 0
	for _, cnt := range counter {
		odd += cnt % 2
	}
	return odd <= k
}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
