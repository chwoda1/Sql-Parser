package sqli

import (
	"fmt"
	"strings"
	"unicode"
)

func Yylex(l *Lexer) error {

	for l.Current() != _eof {

		tok, err := nextToken(l)

		if err != nil {
			panic(err)
		}

		if tok != nil {
			l.tokens <- *tok
		}

		l.Next()

	}

	l.tokens <- *tokenFactory(_eof, nil)
	//close(l.tokens) user can do this

	return nil
}

func nextToken(l *Lexer) (*Token, error) {

	nextChar := l.Current()

	if idStart(nextChar) {
		return tokenizeIdentifier(l)
	} else if isNumeric(nextChar) {
		return tokenizeNumber(l)
	} else if unicode.IsSpace(nextChar) {
		return nil, nil
	}

	switch nextChar {
	case '"':
		return tokenizeDoubleQuote(l)
	case '\'':
		return tokenizeSingleQuote(l)
	case '(':
		return tokenFactory(_lparen, nil), nil
	case ')':
		return tokenFactory(_rparen, nil), nil
	case ',':
		return tokenFactory(_comma, nil), nil
	case '+':
		return tokenFactory(_plus, nil), nil
	case '-':
		{
			// FIXME => Doesn't work for negative numbers
			if unicode.IsDigit(l.Peek()) {
				return tokenizeNumber(l)
			} else {
				return tokenFactory(_minus, nil), nil
			}

		}
	case '*':
		return tokenFactory(_star, nil), nil // Also Kleene Star
	case '/':
		return tokenFactory(_div, nil), nil
	case '%':
		return tokenFactory(_mod, nil), nil
	case '=':
		return tokenFactory(_equal, nil), nil
	case '.':
		return tokenFactory(_dot, nil), nil
	case ';':
		return tokenFactory(_semi, nil), nil
	case '[':
		return tokenFactory(_lbrack, nil), nil
	case ']':
		return tokenFactory(_rbrack, nil), nil
	case '!':
		{
			l.Next()
			switch l.Peek() {
			case '=':
				{
					return tokenFactory(_notEqual, nil), nil
				}
			default:
				{
					return nil, fmt.Errorf("INVALID TOKEN: " + string(nextChar))
				}
			}
		}

	case '<':
		{
			l.Next()
			switch l.Peek() {
			case '=':
				return tokenFactory(_lessEqual, nil), nil
			case '<':
				return tokenFactory(_lShift, nil), nil
			default:
				return tokenFactory(_lessThan, nil), nil
			}
		}
	case '>':
		{
			l.Next()
			switch l.Peek() {
			case '=':
				return tokenFactory(_greaterEqual, nil), nil
			case '>':
				return tokenFactory(_rShift, nil), nil
			default:
				return tokenFactory(_greaterThan, nil), nil

			}
		}

	case ':':
		{
			l.Next()
			switch l.Peek() {
			case ':':
				return tokenFactory(_doubleColon, nil), nil
			default:
				return tokenFactory(_colon, nil), nil
			}
		}
	case '&':
		{
			l.Next()
			switch l.Peek() {
			case '&':
				return tokenFactory(_logicalAnd, nil), nil
			default:
				return tokenFactory(_bitAnd, nil), nil
			}
		}
	case '|':
		{
			l.Next()
			switch l.Peek() {
			case '|':
				return tokenFactory(_logicalOr, nil), nil
			default:
				return tokenFactory(_bitOr, nil), nil
			}
		}

	default:
		return nil, fmt.Errorf("UNSUPPORTED CHARACTER: " + string(nextChar))
	}

	return nil, nil
}

func tokenizeDoubleQuote(l *Lexer) (*Token, error) {

	var sb strings.Builder

	l.Next() // skip double quote that got us to start state

	for {

		nextChar := l.Next()

		switch nextChar {
		case '"':
			return tokenFactory(_string, sb.String()), nil
		case _eof:
			return nil, fmt.Errorf("Unexpected EOF")
		default:
			sb.WriteRune(nextChar)
		}
	}

	return nil, fmt.Errorf("Idk how we got here\n")

}

func tokenizeSingleQuote(l *Lexer) (*Token, error) {

	var sb strings.Builder

	l.Next() // skip first quote

	for {

		nextChar := l.Next()

		switch nextChar {
		case '\'':
			return tokenFactory(_string, sb.String()), nil
		case _eof:
			return nil, fmt.Errorf("Unexpected EOF")
		default:
			sb.WriteRune(nextChar)
		}
	}

	return nil, fmt.Errorf("Idk how we got here\n")

}

// If you want to do anything with this, it needs to be optimized... O(N * K) Lookup =(
// HAT or Burst Radix Trie => http://crpit.com/confpapers/CRPITV62Askitis.pdf?ref=driverlayer.com/web
//
func tokenizeIdentifier(l *Lexer) (*Token, error) {

	// this will hold identifiers
	byteSlice := make([]byte, 0)

	// holds keywords because they need to be standardized like SeLecT
	temp := make([]byte, 0) // optimize this for our custom radix tree implementation because they all suck

	keywords := l.SqlDialect.Keywords()
	root := keywords.Root()

	for {
		x := l.Peek()

		// we have a terminating condition
		if !idStart(x) || x == _eof {

			l.Rewind()

			data, ok := root.Get(temp)

			if ok {
				return tokenFactory(data.(uint), nil), nil
			} else {
				return tokenFactory(_identifier, string(byteSlice)), nil
			}

		}

		byteSlice = append(byteSlice, byte(x))
		temp = append(temp, byte(unicode.ToUpper(x)))
		l.Next()
	}

	return nil, fmt.Errorf("Unexpected EOF\n")

}

func tokenizeNumber(l *Lexer) (*Token, error) {

	var sb strings.Builder

	for {

		next := l.Current()

		if unicode.IsDigit(next) {
			sb.WriteRune(next)
		} else {
			return tokenFactory(_number, sb.String()), nil
		}

		l.Next()
	}
}

func tokenFactory(tokenType uint, tokenValue interface{}) *Token {
	return &Token{tokenType, tokenValue}
}

func isNumeric(token rune) bool {
	if token >= 48 && token <= 57 {
		return true
	}

	return false
}

func idStart(token rune) bool {

	if (token >= 65 && token < 91) || (token >= 97 && token < 123) || token == 95 {
		return true
	}

	return false
}
