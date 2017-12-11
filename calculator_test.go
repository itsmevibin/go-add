package main

import "testing"

func TestAdd(t *testing.T) {

	cases := []struct {
		num1 int
		num2 int
		want int
	}{
		{num1: 1, num2: 3, want: 4},
		{num1: 4, num2: 5, want: 9},
	}

	for _, i := range cases {
		sum := add(i.num1, i.num2)
		if sum != i.want {
			t.Errorf("Expected %d got %d", i.want, sum)
		}
	}
}

func TestSubtract(t *testing.T) {

	cases := []struct {
		num1 int
		num2 int
		want int
	}{
		{num1: 5, num2: 3, want: 2},
		{num1: 10, num2: 21, want: -11},
	}

	for _, i := range cases {
		sub := subtract(i.num1, i.num2)
		if sub != i.want {
			t.Errorf("Subtraction result expected %d but got %d", i.want, sub)
		}
	}
}

func TestMuliply(t *testing.T) {

	cases := []struct {
		num1 int
		num2 int
		want int
	}{
		{num1: 5, num2: 3, want: 15},
		{num1: 10, num2: 21, want: 210},
	}

	for _, i := range cases {
		sub := multipy(i.num1, i.num2)
		if sub != i.want {
			t.Errorf("Multiplication result expected %d but got %d", i.want, sub)
		}
	}
}

func TestDivide(t *testing.T) {

	cases := []struct {
		num1 int
		num2 int
		want int
	}{
		{num1: 15, num2: 3, want: 5},
		{num1: 20, num2: 2, want: 10},
	}

	for _, i := range cases {
		sub := divide(i.num1, i.num2)
		if sub != i.want {
			t.Errorf("Multiplication result expected %d but got %d", i.want, sub)
		}
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		val1     int
		operator string
		val2     int
		want     int
	}{
		{val1: 10, operator: "+", val2: 15, want: 25},
		{val1: 25, operator: "-", val2: 15, want: 10},
		{val1: 20, operator: "/", val2: 5, want: 4},
		{val1: 3, operator: "*", val2: 5, want: 15},
		{val1: 10, operator: "=", want: 10},
	}

	for _, i := range cases {
		result := calculate(i.val1, i.val2, i.operator)
		if result != i.want {
			t.Errorf("Expected (%d%s%d=%d) but got %d", i.val1, i.operator, i.val2, i.want, result)
		}
	}
}
