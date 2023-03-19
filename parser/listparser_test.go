package parser

import (
	"testing"
)

func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	lp := ListParser{}
	s := "[hello world]"
	got := lp.Tokenize(s)
	want := []Token{
		{Type: T_L_Brace, Value: "["},
		{Type: T_ID, Value: "hello"},
		{Type: T_ID, Value: "world"},
		{Type: T_R_Brace, Value: "]"},
	}
	if !sliceEqual(got, want) {
		t.Fatalf("Invalid tokens: wanted: %+v received %+v", got, want)
	}

	s = "[a, b, c]"
	got = lp.Tokenize(s)
	want = []Token{
		{Type: T_L_Brace, Value: "["},
		{Type: T_ID, Value: "a"},
		{Type: T_COMMA, Value: ","},
		{Type: T_ID, Value: "b"},
		{Type: T_COMMA, Value: ","},
		{Type: T_ID, Value: "c"},
		{Type: T_R_Brace, Value: "]"},
	}
	if !sliceEqual(got, want) {
		t.Fatalf("Invalid tokens: wanted: %+v received %+v", got, want)
	}

}
