package winnow

func generateNgram(code []byte) [][]byte {
	ngram := make([][]byte, 0)

	for i := range code {
		if i+N > len(code) {
			break
		}

		ngram = append(ngram, code[i:i+N])
	}

	return ngram
}
