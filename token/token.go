package token

// Type はトークンの種類を表す文字列
type Type string

// Token は`トークンの種類`と`実際のコードの文字列`を持つ
type Token struct {
	Type    Type
	Literal string
}

// Typeの値
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
