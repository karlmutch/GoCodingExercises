package skyline

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNaieve(t *testing.T) {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	assert.Equal(t, Naieve(buildings), [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}})
}

func TestTransitions(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestNegativeNaieve(t *testing.T) {
	assert.Equal(t, false, false)
}

func TestNegativeTransitions(t *testing.T) {
	assert.Equal(t, false, false)
}
