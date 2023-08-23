// For better readability we shouldnt use function over here because its just a waste of time since we can use switch to do it
// But we can use method if we want better organization or making a more advanced calculator
package main

import "fmt"

func main() {
	var choice uint8;
	var value1 float64;
	var value2 float64;
	fmt.Println("================/ Simple Calculator Program /=================")
	fmt.Print("Choose between these options:\n1- Addition\n2- Substract\n3- Multiplication\n4- Division\nUser: ");fmt.Scan(&choice)

	if choice == 1 || choice == 2 || choice == 3 || choice == 4 {
		fmt.Print("\nType your first value\nUser: ");fmt.Scan(&value1);
		fmt.Print("\nType your second value\nUser: ");fmt.Scan(&value2);
	} else {
		fmt.Print("Wrong options!")
	}
	switch choice {
	case 1:
		result := addition(value1, value2); fmt.Printf("The result is %v", result)
	
	case 2:
		result := substract(value1, value2); fmt.Printf("The result is %v", result)
	
	case 3:
		result := multiplication(value1, value2); fmt.Printf("The result is %v", result)
	
	case 4:
		result := division(value1, value2); fmt.Printf("The result is %v", result)	
	}
}

// addition method
func addition(number1 float64, number2 float64) float64 {
	var result float64 = number1 + number2;return result
}

// substraction method
func substract(number1 float64, number2 float64) float64 {
	var result float64 = number1 - number2;return result
}

// multiplication method
func multiplication(number1 float64, number2 float64) float64 {
	var result float64 = number1 * number2;return result
}

// division method
func division(number1 float64, number2 float64) float64 {
	var result float64 = number1 / number2;return result
}
