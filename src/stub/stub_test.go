package stub

import (
	"errors"
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

func TestStrawmanNegativeCase(t *testing.T) {
	word := "Something else"
	foo, err := Strawman(word)
	if err == nil {
		t.Fatal(errors.New("Negative test case failed for strawman"))
	}
	assert.Equal(t, false, foo)
}
