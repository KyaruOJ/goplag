package simtest

import (
	"errors"
	"goplag/lexer"
	"goplag/source"
	"goplag/winnowing"
	"strconv"
)

// Simtest 执行相似度测试
func Simtest(base, srcA, srcB *source.Source, raw bool) (string, error) {
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

	fingerprintsBase := []string{}
	if base != nil {
		tokensBase := lexer.Lexer(base)
		fingerprintsBase = winnowing.Winnowing(tokensBase)
	}

	same := 0
	for _, fpA := range fingerprintsA {
		for _, fpB := range fingerprintsB {
			if fpA == fpB {
				if base != nil {
					flag := false
					for _, fpBase := range fingerprintsBase {
						if fpBase == fpA {
							flag = true
							break
						}
					}

					if !flag {
						same++
					}
				} else {
					same++
				}
				break
			}
		}
	}

	ans := int(same * 100 / len(fingerprintsA))

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
