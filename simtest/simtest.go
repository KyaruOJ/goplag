package simtest

import (
	"errors"
	"goplag/lexer"
	"goplag/source"
	"goplag/winnowing"
)

// Simtest 执行相似度测试
func Simtest(base, srcA, srcB *source.Source) (int, error) {
	if !supportedExt(srcA) || !supportedExt(srcB) {
		return -1, errors.New("Unsupported extension")
	}

	if srcA.Ext != srcB.Ext {
		return -1, errors.New("Extensions are not matched")
	}

	tokensA := lexer.Lexer(srcA)
	tokensB := lexer.Lexer(srcB)

	fingerprintsA := winnowing.Winnowing(tokensA)
	fingerprintsB := winnowing.Winnowing(tokensB)

	fingerprintsBase := []int{}
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

	return ans, nil
}
