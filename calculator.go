package main

import (
	"fmt"
	"os"
	"os/exec"
)

type calcValues struct {
	value    int
	operator string
	result   int
}

func main() {
	c := []calcValues{}
	var result int
	var val int
	var ope string
	result = 0
	fmt.Printf("val : ")
	fmt.Scanf("%d", &result)
	fmt.Printf("ope : ")
	fmt.Scanf("%s", &ope)
	cIntial := calcValues{
		value:    result,
		operator: ope,
		result:   0,
	}
	c = append(c, cIntial)
	for ope != "=" {
		fmt.Printf("val : ")
		fmt.Scanf("%d", &val)
		result = calculate(result, val, ope)
		fmt.Printf("Result : %d\n", result)
		fmt.Printf("ope : ")
		fmt.Scanf("%s", &ope)

		c1 := calcValues{
			value:    val,
			operator: ope,
			result:   result,
		}
		c = append(c, c1)

		if ope == "=" {
			cmd := exec.Command("clear") //Linux example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
			for _, i := range c {
				var operator = i.operator
				if(i.operator == "=") {
					operator = " "+i.operator
				}
				fmt.Printf("%d%s", i.value, operator)
			}
			fmt.Printf(" %v", result)
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
