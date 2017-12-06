package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		vals []int
		want int
	}{
		{[]int{2, 5, 7}, 14},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{}, 00},
	}
	for _, c := range cases {
		sum := add(c.vals)
		if c.want != sum {
			t.Errorf("Sum of arrays %d result is %d but got %d", c.vals, c.want, sum)
		}
	}
}
