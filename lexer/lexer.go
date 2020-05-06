package lexer

import (
	cLexer "goplag/lang/c"
	goLexer "goplag/lang/go"
	javaLexer "goplag/lang/java"
	verilogLexer "goplag/lang/verilog"
	"goplag/source"
	"goplag/token"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Lexer 执行词法分析
func Lexer(src *source.Source) []token.Token {
	input := antlr.NewInputStream(src.Code)

	var lexer antlr.Lexer
	if src.Ext == ".go" {
		lexer = goLexer.NewGoLexer(input)
	} else if src.Ext == ".java" {
		lexer = javaLexer.NewJavaLexer(input)
	} else if src.Ext == ".c" || src.Ext == ".cpp" {
		lexer = cLexer.NewCLexer(input)
	} else if src.Ext == ".v" {
		lexer = verilogLexer.NewSysVerilogHDLLexer(input)
	} else {
		return []token.Token{}
	}
	stream := antlr.NewCommonTokenStream(lexer, 0)
	stream.Fill()

	return token.Create(stream.GetAllTokens(), src.Ext)
}
