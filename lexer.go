package sqli

import (
	"unicode/utf8"
)

type Lexer struct {
	name       string
	input      string
	start      int
	pos        int
	width      int
	tokens     chan Token
	SqlDialect Dialect // SQL Flavor
}

func (l *Lexer) Next() rune {

	var r rune
	if l.pos >= len(l.input) {
		l.width = 0
		return _eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])

	l.pos += l.width
	return r

}

func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Rewind()
	return r
}

func (l *Lexer) Rewind() {
	l.pos -= l.width
}

func (l *Lexer) Ignore() {
	l.pos -= l.width
}

func (l *Lexer) Current() rune {

	if l.pos >= len(l.input) {
		l.width = 0
		return _eof
	}

	r, _ := utf8.DecodeRuneInString(l.input[l.pos:])
	return r
}
