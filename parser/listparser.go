package parser

type TokenType int

const (
	T_L_Brace TokenType = iota
	T_R_Brace
	T_COMMA
	T_ID
)

type Token struct {
	Type  TokenType
	Value string
}

func isChar(x byte) bool {
	return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
}

type ListParser struct{}

func (lp ListParser) Tokenize(s string) []Token {
	b := []byte(s)
	tokens := []Token{}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || b[i] == '\n' || b[i] == '\t' {
			continue
		}
		switch b[i] {
		case ',':
			tokens = append(tokens, Token{Type: T_COMMA, Value: ","})
		case '[':
			tokens = append(tokens, Token{Type: T_L_Brace, Value: "["})
		case ']':
			tokens = append(tokens, Token{Type: T_R_Brace, Value: "]"})
		default:
			if isChar(b[i]) {
				e := i
				for j := i + 1; j < len(b); j++ {
					if isChar(b[j]) {
						e = j
						continue
					} else {
						break
					}
				}
				e++
				tokens = append(tokens, Token{Type: T_ID, Value: string(b[i:e])})
				i = e - 1
				continue
			} else {
				panic("invalid token" + string(b[i]))
			}
		}
	}
	return tokens
}
