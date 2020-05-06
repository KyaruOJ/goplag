package token

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Token 标记
type Token struct {
	Type int
	Text string
}

// Create 创建 token 序列
func Create(tks []antlr.Token, ext string) []Token {
	tokens := []Token{}

	var comment []int
	if ext == ".go" {
		comment = []int{73, 75}
	} else if ext == ".java" {
		comment = []int{109, 110}
	} else if ext == ".c" || ext == ".cpp" {
		comment = []int{117, 118}
	} else if ext == ".v" {
		comment = []int{38, 40}
	}

	for _, tk := range tks {
		commentFlag := false
		for _, c := range comment {
			if tk.GetTokenType() == c {
				commentFlag = true
				break
			}
		}
		if commentFlag {
			continue
		}

		tokens = append(tokens, Token{
			Type: tk.GetTokenType(),
			Text: tk.GetText(),
		})
	}

	return tokens
}
