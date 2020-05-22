package winnowing

import "goplag/token"

// Winnowing 使用 winnowing 算法获得 token 序列的指纹
func Winnowing(tokens []token.Token) map[int]string {
	ngram := generateNgram(tokens)
	hashList := generateHash(ngram)
	fingerprint := generateFingerprint(hashList)

	fps := make(map[int]string)
	for idx, f := range fingerprint {
		fps[idx] = f
	}

	return fps
}
