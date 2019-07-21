package parser

import "fmt"

// Sym are single-rune symbols
type Sym string

// Kinds of single-rune symbols
const (
	SymLParen Sym = "("
	SymRParen     = ")"
	SymLBrace     = "{"
	SymRBrace     = "}"
	SymDot        = "."
	SymComma      = ","
	SymAssign     = "="
)

// Symbol is a token type for single-rune symbols
type Symbol struct {
	sym Sym
	pos Position
}

func (t Symbol) Pos() Position  { return t.pos }
func (t Symbol) String() string { return fmt.Sprintf("Symbol %s", t.sym) }
