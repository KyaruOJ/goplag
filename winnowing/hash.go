package winnowing

import (
	"crypto/sha256"
	"goplag/token"
	"strconv"
)

func generateHash(ngram [][]token.Token) []string {
	hashList := make([]string, 0)

	for _, item := range ngram {
		text := ""
		for _, tk := range item {
			text += strconv.Itoa(tk.Type) + "|"
		}

		hash := getSha256(text)
		hashList = append(hashList, hash)
	}

	return hashList
}

func getSha256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return string(hash[:])
}
