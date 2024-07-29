package lexer

type Lexer struct {
	value        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		value: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.value) {
		l.ch = 0
	} else {
		l.ch = l.value[l.position]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() {

}
