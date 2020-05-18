package winnowing

import "goplag/token"

// Winnowing 使用 winnowing 算法获得 token 序列的指纹
func Winnowing(tokens []token.Token) []string {
	ngram := generateNgram(tokens)
	hashList := generateHash(ngram)
	fingerprint := generateFingerprint(hashList)

	fps := make([]string, 0)
	for _, f := range fingerprint {
		fps = append(fps, f)
	}

	return fps
}
