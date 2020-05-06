package winnowing

import "goplag/token"

func generateNgram(tokens []token.Token) [][]token.Token {
	ngram := make([][]token.Token, 0)

	for i := range tokens {
		if i+N > len(tokens) {
			break
		}

		ngram = append(ngram, tokens[i:i+N])
	}

	return ngram
}
