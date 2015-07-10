package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestStrawman(t *testing.T) {
	word := "strawman"
	foo, err := Strawman(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, foo)
}
