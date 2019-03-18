package sqli

import (
	rtrie "github.com/hashicorp/go-immutable-radix"
	"unicode"
)

type Postgres struct {
	keywords *rtrie.Tree
}

var _KEYWORDS = []struct {
	keyword string
	token   uint
}{

	{"ALTER", _alter}, {"ONLY", _only}, {"SELECT", _select}, {"FROM", _from}, {"WHERE", _where},
	{"LIMIT", _limit}, {"ORDER", _order}, {"GROUP", _group}, {"BY", _by}, {"HAVING", _having},
	{"UNION", _union}, {"ALL", _all}, {"INSERT", _insert}, {"INTO", _into}, {"UPDATE", _update},
	{"DELETE", _delete}, {"IN", _in}, {"IS", _is}, {"NULL", _null}, {"SET", _set},
	{"CREATE", _create}, {"EXTERNAL", _external}, {"TABLE", _table}, {"ASC", _asc}, {"DESC", _desc},
	{"AND", _logicalAnd}, {"OR", _logicalOr}, {"NOT", _not}, {"AS", _as}, {"STORED", _stored},
	{"CSV", _csv}, {"WITH", _with}, {"WITHOUT", _without}, {"ROW", _row}, {"CHAR", _char},
	{"CHARACTER", _character}, {"VARYING", _varying}, {"LARGE", _large}, {"VARCHAR", _varchar}, {"CLOB", _clob},
	{"BINARY", _binary}, {"VARBINARY", _varbinary}, {"BLOB", _blob}, {"FLOAT", _float}, {"REAL", _real},
	{"DOUBLE", _double}, {"PRECISION", _precision}, {"INT", _int}, {"INTEGER", _integer}, {"SMALLINT", _smallint},
	{"BIGINT", _bigint}, {"NUMERIC", _numeric}, {"DECIMAL", _decimal}, {"DEC", _dec}, {"BOOLEAN", _boolean},
	{"DATE", _date}, {"TIME", _time}, {"TIMESTAMP", _timestamp}, {"VALUES", _values}, {"DEFAULT", _default},
	{"ZONE", _zone}, {"REGCLASS", _regclass}, {"TEXT", _text}, {"BYTEA", _bytea}, {"TRUE", _true},
	{"FALSE", _false}, {"COPY", _copy}, {"STDIN", _stdin}, {"PRIMARY", _primary}, {"KEY", _key},
	{"UNIQUE", _unique}, {"UUID", _uuid}, {"ADD", _add}, {"CONSTRAINT", _constraint}, {"FOREIGN", _foreign},
	{"REFERENCES", _references}, {"CASE", _case}, {"WHEN", _when}, {"THEN", _then}, {"ELSE", _else},
	{"END", _end}, {"JOIN", _join}, {"LEFT", _left}, {"RIGHT", _right}, {"FULL", _full}, {"CROSS", _cross},
	{"OUTER", _outer}, {"INNER", _inner}, {"NATURAL", _natural}, {"ON", _on}, {"USING", _using}, {"LIKE", _like}, {"DISTINCT", _distinct},
	{"ILIKE", _ilike},
}

func CreatePostgres() Postgres {
	return Postgres{
		keywords: buildKeywords(),
	}
}

func buildKeywords() *rtrie.Tree {

	tree := rtrie.New()

	for _, x := range _KEYWORDS {
		tree, _, _ = tree.Insert([]byte(x.keyword), x.token)
	}

	return tree
}

// Always pass in a size one array
func (d Postgres) IsIdentifierStart(token rune) bool {

	for _, k := range _KEYWORDS {
		if k.keyword[0] == byte(unicode.ToUpper(token)) {
			return true
		}
	}

	return false
}

func (d Postgres) Keywords() *rtrie.Tree {
	return d.keywords
}
