package token

// Type はトークンの種類を表す文字列
type Type string

// Token はトークンの種類と実際の値を持つ
type Token struct {
	Type    Type
	Literal string
}

// トークンの種類を表す定数
const (
	ILLEGAL = "ILLEGAL" // 解析不能な字句
	EOF     = "EOF"

	// リテラル
	INT = "INT" // 整数 1, 2, 3 など

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ    = "=="
	NOTEQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// ユーザ定義の識別子
	IDENT = "IDENT" // x, y, myVal など

	// 定義済みの識別子（キーワード）
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// 定義済みの識別子とトークンの種類の関連付け
var tokenTypeByKeyword = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupTokenType は、与えられた文字列を基に、トークンの種類を判断して返す
func LookupTokenType(ident string) Type {
	// 与えられた文字列が、定義済みキーワードだった場合(`fn`, `let`など)
	if tokenType, ok := tokenTypeByKeyword[ident]; ok {
		return tokenType
	}
	// 与えられた文字列が、ユーザ定義の識別子だった場合
	return IDENT
}
