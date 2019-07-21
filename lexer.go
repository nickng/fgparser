package parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

// Lexer is a lexer
type Lexer struct {
	r            *bufio.Reader
	ch           chan Token
	lines        []int
	column       int
	lastRuneSize int
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r:  bufio.NewReader(r),
		ch: make(chan Token),
	}
}

func (l *Lexer) read() (rune, error) {
	ch, size, err := l.r.ReadRune()
	if err != nil {
		return ch, err
	}
	if ch == '\n' {
		l.lines = append(l.lines, l.column)
		l.column = 0
	} else {
		l.column += size
	}
	l.lastRuneSize = size
	return ch, nil
}

func (l *Lexer) unread() error {
	err := l.r.UnreadRune()
	if err != nil {
		return err
	}
	if l.column == 0 {
		if len(l.lines) == 0 {
			return errors.New("cannot unread at beginning of file")
		}
		l.column = l.lines[len(l.lines)-1]
		l.lines = l.lines[:len(l.lines)-1]
	} else {
		l.column -= l.lastRuneSize
	}
	return nil
}

func (t *Lexer) Next() (Token, error) {
	ch, err := t.read()
	for {
		if err != nil {
			return nil, err
		}
		if isSpace(ch) {
			if err := t.skipSpace(); err != nil {
				return nil, err
			}
			ch, err = t.read()
		} else {
			break
		}
	}
	pos := pos{line: len(t.lines) + 1, column: t.column}
	switch ch {
	case '(':
		return Symbol{sym: SymLParen, pos: pos}, nil
	case ')':
		return Symbol{sym: SymRParen, pos: pos}, nil
	case '{':
		return Symbol{sym: SymLBrace, pos: pos}, nil
	case '}':
		return Symbol{sym: SymRBrace, pos: pos}, nil
	case '.':
		return Symbol{sym: SymDot, pos: pos}, nil
	case ',':
		return Symbol{sym: SymComma, pos: pos}, nil
	case '=':
		return Symbol{sym: SymAssign, pos: pos}, nil
	case ':':
		return Symbol{sym: SymSemicolon, pos: pos}, nil
	default:
		if err := t.unread(); err != nil {
			return nil, err
		}
		return t.nextIdent()
	}
}

func (t *Lexer) skipSpace() error {
	for {
		ch, err := t.read()
		if err != nil {
			return err
		}
		switch ch {
		case ' ', '\t':
			// skip over
		default:
			t.unread()
			return nil
		}
	}
}

// nextIdent returns the next identifier-type token.
func (t *Lexer) nextIdent() (Token, error) {
	var (
		buf    bytes.Buffer
		line   int
		column int
	)
	line, column = len(t.lines)+1, t.column+1
	for {
		ch, err := t.read()
		if err != nil {
			return nil, err
		}
		if isAlphaNum(ch) {
			buf.WriteRune(ch)
		} else {
			t.unread()
			break
		}
	}
	pos := posRange{line: line, columnStart: column, columnEnd: t.column}
	switch buf.String() {
	case "type":
		return TypeLit{pos: pos}, nil
	case "struct":
		return StructLit{pos: pos}, nil
	case "func":
		return FuncLit{pos: pos}, nil
	case "interface":
		return InterfaceLit{pos: pos}, nil
	case "package":
		return PackageLit{pos: pos}, nil
	case "return":
		return ReturnLit{pos: pos}, nil
	default:
		// A generic identifier
		return Ident{str: buf.String(), pos: pos}, nil
	}
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isAlphaNum(ch rune) bool {
	return ch == '_' ||
		('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		('0' <= ch && ch <= '9')
}
