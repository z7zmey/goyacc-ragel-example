package main

%%{ 
    machine lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type Lexer struct {
	data         []byte
	p, pe, cs    int
	ts, te, act  int

	result []string
}

func NewLexer(data []byte) *Lexer {
    lex := &Lexer{ 
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
    eof := lex.pe
    var tok int

    %%{
        main := |*
            [\t\v\f\r\n ]* => { /* do nothing */ };
            [+(),\-]       => { lval.token = string(lex.data[lex.ts:lex.te]); tok = int(lex.data[lex.ts]); fbreak;};
            [0-9]{2}       => { lval.token = string(lex.data[lex.ts:lex.te]); tok = COUNTRY; fbreak; };
            [0-9]{3}       => { lval.token = string(lex.data[lex.ts:lex.te]); tok = OPERATOR; fbreak; };
            [0-9]{7}       => { lval.token = string(lex.data[lex.ts:lex.te]); tok = PHONE; fbreak; };
        *|;

        write exec;
    }%%

    return int(tok);
}