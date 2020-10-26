package lexer

import (
	"github.com/yakawa/simpleDB/common/helper"
	"github.com/yakawa/simpleDB/common/token"
	"github.com/yakawa/simpleDB/common/value"
)

type lexer struct {
	src        []rune
	currentPos int
	readPos    int
}

func new(src string) *lexer {
	l := &lexer{
		currentPos: 0,
		readPos:    0,
		src:        []rune(src),
	}
	l.readChar()
	return l
}

func Lex(src string) token.Tokens {
	l := new(src)
	tokens := l.lex()

	return tokens
}

func (l *lexer) lex() token.Tokens {
	tokens := l.tokenize()

	return tokens
}

func (l *lexer) readChar() {
	l.currentPos = l.readPos
	l.readPos++
	return
}

func (l *lexer) getCurrentChar() rune {
	if l.currentPos >= len(l.src) {
		return 0
	}
	return l.src[l.currentPos]
}

func (l *lexer) tokenize() token.Tokens {
	tokens := token.Tokens{}

	for {
		ch := l.getCurrentChar()
		if ch == 0 {
			break
		} else if helper.IsWhiteSpace(ch) {
			l.readChar()
		} else {
			t, err := l.findToken()
			tokens = append(tokens, t)
			if err != nil {
				break
			}
		}
	}

	tokens = append(tokens, token.Token{Type: token.EOS})

	return tokens
}

func (l *lexer) findToken() (token.Token, error) {
	ch := l.getCurrentChar()
	switch ch {
	case ';', '+', '-', '*', '/', '%', '(', ')':
		v, tp := l.lookupSymbol()
		t := token.Token{
			Type:    tp,
			Literal: v,
		}
		return t, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		v := l.readNumber()
		val, err := value.Convert(v)
		if err != nil {
			t := token.Token{
				Type:    token.ERROR,
				Literal: v,
			}
			return t, err
		}
		t := token.Token{
			Literal: v,
			Value:   val,
			Type:    token.NUMBER,
		}
		return t, nil
	default:
		v := l.readIdent()
		isKeyword, tp := token.CheckKeyword(v)
		t := token.Token{
			Literal: v,
		}
		if isKeyword {
			t.Type = tp
		} else {
			t.Type = token.IDENT
		}
		return t, nil
	}
}

func (l *lexer) readIdent() string {
	v := []rune("")
	for {
		ch := l.getCurrentChar()
		if ch == 0 {
			break
		} else if helper.IsWhiteSpace(ch) {
			break
		} else if helper.IsSymbol(ch) {
			break
		}
		v = append(v, ch)
		l.readChar()
	}
	return string(v)
}

func (l *lexer) readNumber() string {
	v := []rune("")
	for {
		ch := l.getCurrentChar()
		if helper.IsDigit(ch) {
			v = append(v, ch)
			l.readChar()
			continue
		} else {
			break
		}
	}
	return string(v)
}

func (l *lexer) lookupSymbol() (string, token.Type) {
	ch := l.getCurrentChar()
	var v string
	var val token.Type
	switch ch {
	case ';':
		val = token.S_SEMICOLON
		v = ";"
	case '+':
		val = token.S_PLUS
		v = "+"
	case '-':
		val = token.S_MINUS
		v = "-"
	case '*':
		val = token.S_ASTERISK
		v = "*"
	case '/':
		val = token.S_SOLIDAS
		v = "/"
	case '%':
		val = token.S_PERCENT
		v = "%"
	case '(':
		val = token.S_LPAREN
		v = "("
	case ')':
		val = token.S_RPAREN
		v = ")"
	default:
		val = token.UNKNOWN
		v = ""
	}
	l.readChar()
	return v, val
}
