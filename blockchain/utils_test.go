package blockchain

import (
	"fmt"
	"testing"
)

func TestCalculateBlockReward(t *testing.T) {
	cases := []struct {
		height uint32
		value  int64
	}{
		{1, 5000000000},
		{210000, 2500000000},
		{420000, 1250000000},
		{630000, 625000000},
		{840000, 312500000},
		{1050000, 156250000},
		{1260000, 78125000},
		{1470000, 39062500},
		{2100000, 4882812},
		{2940000, 305175},
		{3150000, 152587},
		{3360000, 76293},
		{4410000, 2384},
		{5460000, 74},
		{6300000, 4},
		{6720000, 1},
		{6930000, 0},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprint(tc.height), func(t *testing.T) {
			actual := CalculateBlockReward(tc.height)
			expected := tc.value

			if actual != expected {
				t.Errorf("Expected %d to equal %d", actual, expected)
			}
		})
	}
}

func TestCalculateHash(t *testing.T) {
	cases := []struct {
		key, hash string
	}{
		{"Hello world!", "c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a"},
		{fmt.Sprint(Block{Height: 1}), "bfbe3184a287068b05958105fa0ec0a167609d5c850e864d5086cfc50aa28a3e"},
	}

	for _, tc := range cases {
		t.Run(tc.key, func(t *testing.T) {
			actual := fmt.Sprintf("%x", CalculateHash(tc.key))
			expected := tc.hash

			if actual != expected {
				t.Errorf("Expected %s to equal %s", actual, expected)
			}
		})
	}
}
