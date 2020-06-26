package simtest

import (
	"fmt"
	"goplag/lexer"
	"goplag/source"
	"goplag/winnowing"
	"sort"
	"strconv"
)

// Simtest 执行相似度测试
func Simtest(srcs []*source.Source) error {
	for _, src := range srcs {
		src.Tokens = lexer.Lexer(src)
		src.Fingerprints = winnowing.Winnowing(src.Tokens)
	}

	for _, srcA := range srcs {
		for _, srcB := range srcs {
			if srcA == srcB {
				continue
			}

			ans := Compare(srcA, srcB)

			fmt.Println(srcA.Path + "|" + srcB.Path + "|" + strconv.Itoa(ans))
		}
	}

	return nil
}

// Compare 执行对比测试
func Compare(a, b *source.Source) int {
	interval := make([]int, 0)

	for idx, fpA := range a.Fingerprints {
		for _, fpB := range b.Fingerprints {
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

		ans = int(length * 100 / len(a.Tokens))
	}

	return ans
}
