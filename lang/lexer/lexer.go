
package lexer

import (
	
	// "fmt"
	// "github.com/awalterschulze/picolang/lang/util"
	
	"io/ioutil"
	"unicode/utf8"
	"github.com/awalterschulze/picolang/lang/token"
)

const(
	NoState = -1
	NumStates = 119
	NumSymbols = 83
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '-'
1: '.'
2: '.'
3: '.'
4: 'l'
5: 'e'
6: 't'
7: '='
8: 'i'
9: 'm'
10: 'p'
11: 'o'
12: 'r'
13: 't'
14: 'p'
15: 'r'
16: 'i'
17: 'n'
18: 't'
19: 'l'
20: 'n'
21: '('
22: ')'
23: '('
24: ','
25: '{'
26: '}'
27: '/'
28: '/'
29: '\n'
30: '/'
31: '*'
32: '*'
33: '*'
34: '/'
35: ' '
36: '\t'
37: '\n'
38: '\r'
39: '0'
40: '0'
41: 'x'
42: 'X'
43: 'e'
44: 'E'
45: '+'
46: '-'
47: '_'
48: '_'
49: '.'
50: '\'
51: 'U'
52: '\'
53: 'u'
54: '\'
55: 'x'
56: '\'
57: '`'
58: '`'
59: '\'
60: 'a'
61: 'b'
62: 'f'
63: 'n'
64: 'r'
65: 't'
66: 'v'
67: '\'
68: '''
69: '"'
70: '"'
71: '"'
72: '''
73: '''
74: '0'-'9'
75: '0'-'7'
76: '0'-'9'
77: 'A'-'F'
78: 'a'-'f'
79: '1'-'9'
80: 'A'-'Z'
81: 'a'-'z'
82: .

*/
