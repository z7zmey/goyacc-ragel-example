
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 5
const lexer_first_final int = 5
const lexer_error int = 0

const lexer_en_main int = 5


//line lexer.rl:9


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
    
//line lexer.go:32
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:25
    return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
    eof := lex.pe
    var tok int

    
//line lexer.go:53
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 5:
		goto st_case_5
	case 0:
		goto st_case_0
	case 6:
		goto st_case_6
	case 1:
		goto st_case_1
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	}
	goto st_out
tr2:
//line lexer.rl:41
( lex.p) = ( lex.te) - 1
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = OPERATOR; {( lex.p)++;  lex.cs = 5; goto _out } }
	goto st5
tr5:
//line lexer.rl:42
 lex.te = ( lex.p)+1
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = PHONE; {( lex.p)++;  lex.cs = 5; goto _out } }
	goto st5
tr7:
//line lexer.rl:39
 lex.te = ( lex.p)+1
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 5; goto _out }}
	goto st5
tr9:
//line lexer.rl:38
 lex.te = ( lex.p)
( lex.p)--
{ /* do nothing */ }
	goto st5
tr10:
//line lexer.rl:40
 lex.te = ( lex.p)
( lex.p)--
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = COUNTRY; {( lex.p)++;  lex.cs = 5; goto _out } }
	goto st5
tr12:
//line lexer.rl:41
 lex.te = ( lex.p)
( lex.p)--
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = OPERATOR; {( lex.p)++;  lex.cs = 5; goto _out } }
	goto st5
	st5:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:123
		if  lex.data[( lex.p)] == 32 {
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 40:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st6
			}
		case  lex.data[( lex.p)] > 41:
			switch {
			case  lex.data[( lex.p)] > 45:
				if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
					goto st1
				}
			case  lex.data[( lex.p)] >= 43:
				goto tr7
			}
		default:
			goto tr7
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if  lex.data[( lex.p)] == 32 {
			goto st6
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
			goto st6
		}
		goto tr9
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto st0
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr11
		}
		goto tr10
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st8
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer.go:189
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st2
		}
		goto tr12
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st3
		}
		goto tr2
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st4
		}
		goto tr2
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr5
		}
		goto tr2
	st_out:
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 6:
			goto tr9
		case 7:
			goto tr10
		case 8:
			goto tr12
		case 2:
			goto tr2
		case 3:
			goto tr2
		case 4:
			goto tr2
		}
	}

	_out: {}
	}

//line lexer.rl:46


    return int(tok);
}