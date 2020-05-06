package winnowing

import (
	"goplag/token"
	"math"
)

func generateHash(ngram [][]token.Token) []int {
	hashList := make([]int, 0)
	h := 0

	for i := 0; i < N; i++ {
		h += int(ngram[0][i].Type) * int(math.Pow(float64(Base), float64(N-1-i)))
	}

	hashList = append(hashList, h)

	for i := 1; i < len(ngram); i++ {
		h = h*Base - int(ngram[i-1][0].Type)*int(math.Pow(float64(Base), float64(N))) + int(ngram[i][N-1].Type)
		hashList = append(hashList, h)
	}

	return hashList
}
