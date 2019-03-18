package sqli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testMath = &Lexer{
	name: "testLexer",
	input: "SELECT relname, 100 * idx_scan / (seq_scan + idx_scan) percent_of_times_index_used, n_live_tup rows_in_table " +
		"FROM pg_stat_user_tables " +
		"ORDER BY n_live_tup DESC",
	start:      0,
	pos:        0,
	width:      1,
	tokens:     make(chan Token),
	SqlDialect: CreatePostgres(),
}

var mathResult = []struct {
	token uint
	input string
}{
	{uint(_select), "SELECT"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_comma), "COMMA"},
	{uint(_number), "NUMBER"},
	{uint(_star), "STAR"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_div), "DIVIDE"},
	{uint(_lparen), "LPAREN"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_plus), "PLUS"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_rparen), "RPAREN"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_comma), "COMMA"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_identifier), "IDENTIFIER"},

	// Line 2
	{uint(_from), "FROM"},
	{uint(_identifier), "IDENTIFIER"},

	// Line 3
	{uint(_order), "ORDER"},
	{uint(_by), "BY"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_desc), "DESC"},
}

func TestMathLexer(t *testing.T) {

	result := make([]Token, 0)

	go func() {
		err := Yylex(testMath)

		if err != nil {
			panic(err)
		}
	}()

	for {
		x := <-testMath.tokens

		if x.Type == _eof {
			break
		}
		result = append(result, x)

	}

	for i := 0; i < len(result); i++ {
		t.Run(mathResult[i].input, func(t *testing.T) {
			assert.Equal(t, mathResult[i].token, result[i].Type)
		})
	}
}

var testFunctionQuery = &Lexer{
	name: "testLexer",
	input: "SELECT pid, age(clock_timestamp(), query_start), usename, query " +
		"FROM pg_stat_activity " +
		"WHERE query != '<IDLE>' AND query NOT ILIKE '%pg_stat_activity%' " +
		"ORDER BY query_start desc",
	start:      0,
	pos:        0,
	width:      1,
	tokens:     make(chan Token),
	SqlDialect: CreatePostgres(),
}

var functionOut = []struct {
	token uint
	input string
}{
	// LINE 1
	{uint(_select), "SELECT"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_comma), "COMMA"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_lparen), "LPAREN"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_lparen), "LPAREN"},
	{uint(_rparen), "RPAREN"},
	{uint(_comma), "COMMA"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_rparen), "RPAREN"},
	{uint(_comma), "COMMA"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_comma), "COMMA"},
	{uint(_identifier), "IDENTIFIER"},

	// LINE 2
	{uint(_from), "FROM"},
	{uint(_identifier), "IDENTIFIER"},

	// LINE 3
	{uint(_where), "WHERE"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_notEqual), "NOT EQUAL"},
	{uint(_string), "STRING"},
	{uint(_logicalAnd), "AND"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_not), "NOT"},
	{uint(_ilike), "ILIKE"},
	{uint(_string), "STRING"},

	// LINE 4
	{uint(_order), "ORDER"},
	{uint(_by), "BY"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_desc), "DESC"},
	{uint(_semi), "SEMICOLON"},
}

func TestFunctionAndString(t *testing.T) {

	result := make([]Token, 0)

	go func() {
		err := Yylex(testFunctionQuery)

		if err != nil {
			panic(err)
		}
	}()

	for {
		x := <-testFunctionQuery.tokens

		if x.Type == _eof {
			break
		}

		result = append(result, x)
	}

	for i := 0; i < len(result); i++ {
		t.Run(functionOut[i].input, func(t *testing.T) {
			assert.Equal(
				t, functionOut[i].token, result[i].Type)
		})
	}
}

var testOperatorQuery = &Lexer{
	name: "testLexer",
	input: "SELECT * FROM asdf " +
		"WHERE 5 & 3 = 1 " +
		"AND 5 >= 2 " +
		"&& 5 > 2 " +
		"|| 5 <= 2 OR 1 | 1 = 1",
	start:      0,
	pos:        0,
	width:      1,
	tokens:     make(chan Token),
	SqlDialect: CreatePostgres(),
}

var operatorResult = []struct {
	token uint
	input string
}{
	{uint(_select), "SELECT"},
	{uint(_star), "STAR"},
	{uint(_from), "FROM"},
	{uint(_identifier), "IDENTIFIER"},
	{uint(_where), "WHERE"},
	{uint(_number), "NUMBER"},
	{uint(_bitAnd), "BIT AND"},
	{uint(_number), "NUMBER"},
	{uint(_equal), "EQUAL"},
	{uint(_number), "NUMBER"},
	{uint(_logicalAnd), "LOGICAL AND"},
	{uint(_number), "NUMBER"},
	{uint(_greaterEqual), "GREATER EQUAL"},
	{uint(_number), "NUMBER"},
	{uint(_logicalAnd), "LOGICAL AND"},
	{uint(_number), "NUMBER"},
	{uint(_greaterThan), "GREATER"},
	{uint(_number), "NUMBER"},
	{uint(_logicalOr), "LOGICAL OR"},
	{uint(_number), "NUMBER"},
	{uint(_lessEqual), "LESS EQUAL"},
	{uint(_number), "NUMBER"},
	{uint(_logicalOr), "LOGICAL OR"},
	{uint(_number), "NUMBER"},
	{uint(_bitOr), "BIT OR"},
	{uint(_number), "NUMBER"},
	{uint(_equal), "EQUAL"},
	{uint(_number), "NUMBER"},
}

func TestOperatorQuery(t *testing.T) {

	result := make([]Token, 0)

	go func() {
		err := Yylex(testOperatorQuery)

		if err != nil {
			panic(err)
		}
	}()

	for {
		x := <-testOperatorQuery.tokens

		if x.Type == _eof {
			break
		}

		result = append(result, x)
	}

	for i := 0; i < len(result); i++ {
		t.Run(operatorResult[i].input, func(t *testing.T) {
			assert.Equal(t, operatorResult[i].token, result[i].Type)
		})
	}
}

// Benchmarks
var benchOperatorQuery = &Lexer{
	name: "testLexer",
	input: "SELECT * FROM asdf " +
		"WHERE 5 & 3 = 1 " +
		"AND 5 >= 2 " +
		"&& 5 > 2 " +
		"|| 5 <= 2 OR 1 | 1 = 1",
	start:      0,
	pos:        0,
	width:      1,
	tokens:     make(chan Token),
	SqlDialect: CreatePostgres(),
}

func BenchmarkSqlParse(b *testing.B) {

	b.ReportAllocs()

	for n := 0; n < b.N; n++ {

		result := make([]Token, 0)

		go func() {
			err := Yylex(benchOperatorQuery)

			if err != nil {
				panic(err)
			}
		}()

		for {
			x := <-benchOperatorQuery.tokens

			if x.Type == _eof {
				break
			}

			result = append(result, x)
		}
	}

}

// Benchmarks

var benchFunctionQuery = &Lexer{
	name: "testLexer",
	input: "SELECT pid, age(clock_timestamp(), query_start), usename, query " +
		"FROM pg_stat_activity " +
		"WHERE query != '<IDLE>' AND query NOT ILIKE '%pg_stat_activity%' " +
		"ORDER BY query_start desc",
	start:      0,
	pos:        0,
	width:      1,
	tokens:     make(chan Token),
	SqlDialect: CreatePostgres(),
}

func xBenchmarkFunctionString(b *testing.B) {

	b.ReportAllocs()

	for n := 0; n < b.N; n++ {

		result := make([]Token, 0)
		go func() {
			err := Yylex(benchFunctionQuery)

			if err != nil {
				panic(err)
			}
		}()

		for {
			x := <-benchFunctionQuery.tokens

			if x.Type == _eof {
				break
			}

			result = append(result, x)
		}

	}

}
