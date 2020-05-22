package simtest

import (
	"errors"
	"goplag/lexer"
	"goplag/source"
	"goplag/winnowing"
	"strconv"
)

// Simtest 执行相似度测试
func Simtest(srcA, srcB *source.Source, raw bool) (string, error) {
	if !supportedExt(srcA) || !supportedExt(srcB) {
		return "", errors.New("Unsupported extension")
	}

	if srcA.Ext != srcB.Ext {
		return "", errors.New("Extensions are not matched")
	}

	tokensA := lexer.Lexer(srcA)
	tokensB := lexer.Lexer(srcB)

	fingerprintsA := winnowing.Winnowing(tokensA)
	fingerprintsB := winnowing.Winnowing(tokensB)
	fingerprintsSame := make(map[int]string)

	same := 0
	for idx, fpA := range fingerprintsA {
		for _, fpB := range fingerprintsB {
			if fpA == fpB {
				fingerprintsSame[idx] = fpA
				break
			}
		}
	}

	ans := int(len(fingerprintsSame) * 100 / len(fingerprintsA))

	if raw {
		res := ""

		res += "similarity: "
		res += strconv.Itoa(ans) + "% "
		res += "(" + strconv.Itoa(same) + " / " + strconv.Itoa(len(fingerprintsA)) + ")" + "\n"

		res += "tokens: "
		res += strconv.Itoa(len(tokensA)) + " - " + strconv.Itoa(len(tokensB)) + "\n"

		res += "fingerprints: "
		res += strconv.Itoa(len(fingerprintsA)) + " - " + strconv.Itoa(len(fingerprintsB))

		return res, nil
	}

	return strconv.Itoa(ans), nil
}
