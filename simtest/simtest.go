package simtest

import (
	"errors"
	"goplag/lexer"
	"goplag/source"
	"goplag/winnow"
	"goplag/winnowing"
	"sort"
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
	interval := make([]int, 0)

	for idx, fpA := range fingerprintsA {
		for _, fpB := range fingerprintsB {
			if fpA == fpB {
				interval = append(interval, idx)
				break
			}
		}
	}

	ans := 0
	length := 0

	if len(interval) > 0 {
		sort.Ints(interval)
		N := winnowing.N
		merged := make([][]int, 0)
		merged = append(merged, []int{interval[0], interval[0] + N})
		for _, startAt := range interval {
			m := merged[len(merged)-1]

			if startAt > m[1] {
				merged = append(merged, []int{startAt, startAt + N})
				continue
			}

			if startAt+N > m[1] {
				m[1] = startAt + N
			}
		}

		length = 0
		for _, item := range merged {
			length += item[1] - item[0]
		}

		ans = int(length * 100 / len(tokensA))
	}

	if raw {
		res := ""

		res += "similarity: "
		res += strconv.Itoa(ans) + "% "
		res += "(" + strconv.Itoa(length) + " / " + strconv.Itoa(len(tokensA)) + ")" + "\n"

		res += "tokens: "
		res += strconv.Itoa(len(tokensA)) + " - " + strconv.Itoa(len(tokensB))

		return res, nil
	}

	return strconv.Itoa(ans), nil
}

// Origtest 执行原始相似度测试
func Origtest(srcA, srcB *source.Source) (string, error) {
	if !supportedExt(srcA) || !supportedExt(srcB) {
		return "", errors.New("Unsupported extension")
	}

	if srcA.Ext != srcB.Ext {
		return "", errors.New("Extensions are not matched")
	}

	tokensA := lexer.Lexer(srcA)
	tokensB := lexer.Lexer(srcB)

	fingerprintsA := winnow.Winnowing(tokensA)
	fingerprintsB := winnow.Winnowing(tokensB)

	same := 0
	for _, fpA := range fingerprintsA {
		for _, fpB := range fingerprintsB {
			if fpA == fpB {
				same++
				break
			}
		}
	}

	ans := int(same * 100 / len(fingerprintsA))

	return strconv.Itoa(ans), nil
}
