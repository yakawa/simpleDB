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
	case ';', '+', '-', '*', '/', '%':
		v, val := l.lookupSymbol()
		t := token.Token{
			Type:    token.SYMBOL,
			Literal: v,
			Value:   val,
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
		val, err := value.Convert(v)
		isKeyword, _ := value.CheckKeyword(v)
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
		}
		if isKeyword {
			t.Type = token.KEYWORD
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

func (l *lexer) lookupSymbol() (string, value.Value) {
	ch := l.getCurrentChar()
	var v string
	val := value.Value{}
	switch ch {
	case ';':
		val.Type = value.S_SEMICOLON
		v = ";"
	case '+':
		val.Type = value.S_PLUS
		v = "+"
	case '-':
		val.Type = value.S_MINUS
		v = "-"
	case '*':
		val.Type = value.S_ASTERISK
		v = "*"
	case '/':
		val.Type = value.S_SOLIDAS
		v = "/"
	case '%':
		val.Type = value.S_PERCENT
		v = "%"
	default:
		val.Type = value.UNKNOWN
		v = ""
	}
	l.readChar()
	return v, val
}
