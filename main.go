package main

//go:generate goyacc -o parser.go parser.y
//go:generate ragel -Z -G2 -o lexer.go lexer.rl

func main() {
	lexer := NewLexer([]byte(`+38(093)1234567, 38 098 7654321`))
	yyParse(lexer)

	for k, v := range lexer.result {
		println(k, v)
	}
}