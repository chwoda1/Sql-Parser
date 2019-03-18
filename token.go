package sqli

import (
	rtrie "github.com/hashicorp/go-immutable-radix"
)

type Dialect interface {
	Keywords() *rtrie.Tree
	IsIdentifierStart(rune) bool
}

type Token struct {
	Type uint // FIXME
	Val  interface{}
}

const (
	_eof           = iota
	_select        = iota
	_drop          = iota
	_where         = iota
	_between       = iota
	_logicalAnd    = iota // AND or &&
	_logicalOr     = iota // OR  or ||
	_equal         = iota // IS  or =
	_not           = iota
	_in            = iota
	_nullSafeEqual = iota
	_greaterThan   = iota
	_greaterEqual  = iota
	_lessThan      = iota
	_lessEqual     = iota
	_like          = iota
	_notEqual      = iota
	_plus          = iota
	_minus         = iota
	_bitOr         = iota // |
	_bitAnd        = iota
	_lShift        = iota
	_rShift        = iota
	_div           = iota // DIV or /
	_mod           = iota // MOD or %
	_xor           = iota // XOR or ^
	_lparen        = iota
	_rparen        = iota
	_lbrack        = iota
	_rbrack        = iota
	_semi          = iota
	_comma         = iota
	_atat          = iota // @@ => Global Var IMPLEMENT FIXME
	_star          = iota // Kleene Closure
	_dot           = iota
	_pound         = iota // comment start
	_fslash        = iota // comment start
	_colon         = iota
	_doubleColon   = iota // IDK
	_union         = iota
	_all           = iota
	_from          = iota
	_string        = iota
	_identifier    = iota
	_number        = iota
	_bool          = iota
	_alter         = iota
	_only          = iota
	_limit         = iota
	_order         = iota
	_group         = iota
	_by            = iota
	_having        = iota
	_insert        = iota
	_into          = iota
	_update        = iota
	_delete        = iota
	_is            = iota
	_null          = iota
	_set           = iota
	_create        = iota
	_external      = iota
	_table         = iota
	_asc           = iota
	_desc          = iota
	_stored        = iota
	_csv           = iota
	_with          = iota
	_without       = iota
	_row           = iota
	_char          = iota
	_character     = iota
	_varying       = iota
	_large         = iota
	_varchar       = iota
	_clob          = iota
	_binary        = iota
	_varbinary     = iota
	_blob          = iota
	_float         = iota
	_real          = iota
	_double        = iota
	_precision     = iota
	_int           = iota
	_integer       = iota
	_smallint      = iota
	_bigint        = iota
	_numeric       = iota
	_decimal       = iota
	_dec           = iota
	_boolean       = iota
	_date          = iota
	_time          = iota
	_timestamp     = iota
	_values        = iota
	_default       = iota
	_zone          = iota
	_regclass      = iota
	_text          = iota
	_bytea         = iota
	_true          = iota
	_false         = iota
	_copy          = iota
	_stdin         = iota
	_primary       = iota
	_key           = iota
	_unique        = iota
	_uuid          = iota
	_add           = iota
	_constraint    = iota
	_foreign       = iota
	_references    = iota
	_case          = iota
	_when          = iota
	_then          = iota
	_else          = iota
	_end           = iota
	_join          = iota
	_left          = iota
	_right         = iota
	_full          = iota
	_cross         = iota
	_outer         = iota
	_inner         = iota
	_natural       = iota
	_on            = iota
	_using         = iota
	_as            = iota
	_distinct      = iota
	_ilike         = iota
)
