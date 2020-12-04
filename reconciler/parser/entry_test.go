package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimalValidEntry(t *testing.T) {
	yaml := `
date: 2020-01-01
`
	e, err := Parse(yaml)
	assert.Equal(t, e.Summary(), "")
	assert.Equal(t, err, nil)
}

func TestFailOnUnknownProperties(t *testing.T) {
	yaml := `
date: 2020-01-01
foo: 1
bar: test
`
	e, err := Parse(yaml)
	assert.Equal(t, e, nil)
	assert.Error(t, err)
}
