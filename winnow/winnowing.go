package winnow

import "goplag/token"

// Winnowing 使用 winnowing 算法获得 token 序列的指纹
func Winnowing(tokens []token.Token) map[int]int {
	ngram := generateNgram(tokens)
	hashList := generateHash(ngram)
	fingerprint := generateFingerprint(hashList)

	fps := make(map[int]int)
	for idx, f := range fingerprint {
		fps[idx] = f
	}

	return fps
}
