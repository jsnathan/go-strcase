// Copyright (c) 2017, A. Stoewer <adrian.jsnathan@rz.ifi.lmu.de>
// All rights reserved.

package strcase_test

import (
	"testing"

	"github.com/jsnathan/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "f",
		"Foo":                        "foo",
		"FooB":                       "foo_b",
		"FooID":                      "foo_id",
		" FooBar\t":                  "foo_bar",
		"HTTPStatusCode":             "http_status_code",
		"ParseURL.DoParse":           "parse_url.do_parse",
		"Convert Space":              "convert_space",
		"Convert-dash":               "convert_dash",
		"Skip___MultipleUnderscores": "skip_multiple_underscores",
		"Skip   MultipleSpaces":      "skip_multiple_spaces",
		"Skip---MultipleDashes":      "skip_multiple_dashes",
	}

	for camel, snake := range data {
		converted := strcase.SnakeCase(camel)
		assert.Equal(t, snake, converted)
	}
}
