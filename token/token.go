package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 解析不能な字句
	EOF     = "EOF"

	// 識別子
	IDENT = "IDENT" // x, y など

	// リテラル
	INT = "INT" // 整数 1, 2, 3 など

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
