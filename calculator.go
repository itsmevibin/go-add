package main

import (
	"fmt"
)

type calcValues struct {
	value    int
	operator string
	result   int
}

func main() {
	c := [10]calcValues{}
	var result int
	var val int
	var ope string
	result = 0
	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Printf("val : ")
			fmt.Scanf("%d", &result)
			fmt.Printf("ope : ")
			fmt.Scanf("%s", &ope)
		}

		fmt.Printf("val : ")
		fmt.Scanf("%d", &val)
		if i > 0 {
			result = c[i-1].result
		}
		result = calculate(result, val, ope)

		fmt.Printf("Result : %d\n", result)
		fmt.Printf("ope : ")
		fmt.Scanf("%s", &ope)
		c1 := calcValues{
			value:    val,
			operator: ope,
			result:   result,
		}
		c[i] = c1
		if ope == "=" {
			for _, i := range c {
				fmt.Printf("%d%s", i.value, i.operator)
			}
			fmt.Printf("%d", result)
			break
		}
	}
}

func calculate(lastResult int, val int, operator string) (result int) {
	switch operator {
	case "=":
		result = lastResult
	case "+":
		result = add(lastResult, val)
	case "-":
		result = subtract(lastResult, val)
	case "*":
		result = multipy(lastResult, val)
	case "/":
		result = divide(lastResult, val)
	}
	return
}

func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func subtract(num1 int, num2 int) (sub int) {
	sub = num1 - num2
	return
}

func multipy(num1 int, num2 int) (sub int) {
	sub = num1 * num2
	return
}

func divide(num1 int, num2 int) (sub int) {
	sub = num1 / num2
	return
}
