package main

func calculate(condition string, values ...int) int {
	if condition == "multiply" {
		multiply := 1
		for _, number := range values {
			multiply *= number
		}
		return multiply

	} else if condition == "add" {
		add := 0
		for _, number := range values {
			add += number
		}
		return add
	} else {
		subtract := 0
		for _, number := range values {
			subtract -= number
		}
		return subtract
	}

}
