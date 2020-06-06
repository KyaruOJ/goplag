package winnow

import (
	"math"
)

func generateHash(ngram [][]byte) []int {
	hashList := make([]int, 0)
	h := 0

	for i := 0; i < N; i++ {
		h += int(ngram[0][i]) * int(math.Pow(float64(Base), float64(N-1-i)))
	}

	hashList = append(hashList, h)

	for i := 1; i < len(ngram); i++ {
		h = h*Base - int(ngram[i-1][0])*int(math.Pow(float64(Base), float64(N))) + int(ngram[i][N-1])
		hashList = append(hashList, h)
	}

	return hashList
}
