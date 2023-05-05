package did

import "errors"

type data struct {
	val  string
	desc string
	err  error
}

var prefixes = []data{
	{"ab", "two alpha chars", nil},
	{"abc", "three alpha chars", nil},
	{"AZ", "uppercase", nil},
	{"AbC", "upper and lowercase combined", nil},
	{"abcd", "too many chars", errors.New("")},
	{"a", "not enough chars", errors.New("")},
	{"", "empty", errors.New("")},
	{".$", "has special chars", errors.New("")},
	{"k9", "has numbers", errors.New("")},
}

var separators = []data{
	{"_", "underscore", nil},
	{"-", "hyphen", nil},
	{"+", "plus sign", nil},
	{"", "empty/ no separator", nil},
	{"%", "invalid symbol", errors.New("")},
	{"--", "more than one character", errors.New("")},
	{" ", "whitespace", errors.New("")},
	{"^", "caret (circumflex)", errors.New("")},
}

var invalidStrs = map[string]string{
	"ab-526cac35b-e74429beb4f2ecca5-6c57":  "more than 1 separator",
	"a9-526cac35b7e74429beb4f2ecca56c571":  "prefix invalid",
	"ab=526cac35b7e74429beb4f2ecca56c571":  "separator invalid",
	"ab-526cac357e74429beb4f2ecca56c571":   "hex has not enough chars",
	"ab-526cac35b7e74429beb4f2ecca56c5711": "hex has too many chars",
}
