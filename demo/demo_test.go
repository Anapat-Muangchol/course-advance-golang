package demo_test

import (
	"demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstCase(t *testing.T) {
	expected := "Hello"
	r, _ := demo.SayHi()
	assert.Equal(t, r, expected, "they should be equal")
}
