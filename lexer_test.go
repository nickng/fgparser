package parser

import (
	"io"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	lexer := NewLexer(strings.NewReader("func x (v type) nil {} more()"))
	var tokens []Token
	for {
		tok, err := lexer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		tokens = append(tokens, tok)
	}
	expected := []Token{
		FuncLit{},
		Ident{str: "x"},
		Symbol{sym: "("},
		Ident{str: "v"},
		TypeLit{},
		Symbol{sym: ")"},
		Ident{str: "nil"},
		Symbol{sym: "{"},
		Symbol{sym: "}"},
		Ident{str: "more"},
		Symbol{sym: "("},
		Symbol{sym: ")"},
	}
	if want, got := len(expected), len(tokens); want != got {
		t.Fatalf("expect length of tokens slice to be the same (want %d got %d)", want, got)
	}
	for i := range expected {
		t.Logf("%s == %s", expected[i].String(), tokens[i].String())
		if want, got := expected[i].String(), tokens[i].String(); want != got {
			t.Fatalf("parsed tokens do not match: %s != %s", want, got)
		}
	}
}
