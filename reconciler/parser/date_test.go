package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbsentDatePropertyFails(t *testing.T) {
	yaml := `
summary: Just a normal day
`
	e, err := Parse(yaml)
	assert.Equal(t, e, nil)
	assert.Error(t, err)
}

func TestMalformedDateFails(t *testing.T) {
	yaml := `
date: 01.01.2020
`
	e, err := Parse(yaml)
	assert.Equal(t, e, nil)
	assert.Error(t, err)
}
