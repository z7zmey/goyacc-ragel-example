%{
package main

%}

%union{
    phone    string
    token    string
}

%token <token> COUNTRY
%token <token> OPERATOR
%token <token> PHONE

%type <phone> phone

%%
start:
        phone {
            yylex.(*Lexer).result = []string{$1}
        }
    |   start ',' phone {
            yylex.(*Lexer).result = append(yylex.(*Lexer).result, $3)
        }

phone:
        '+' COUNTRY '(' OPERATOR ')' PHONE {
            $$ = "+"+$2+"("+$4+")"+$6
        }
    |   COUNTRY OPERATOR PHONE {
            $$ = "+"+$1+"("+$2+")"+$3
        }
%%
