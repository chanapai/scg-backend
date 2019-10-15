package main

import "fmt"

func getSum(num1 int) int {
	a := (num1+1)*num1 + 3

	return a
}

func main() {
	var text string
	char := []string{"x", "5", "9", "15", "23", "y", "z"}
	fmt.Println("X, 5, 9, 15, 23, Y, Z")
	fmt.Println("Type x, y, or z to get value >>")
	fmt.Scanf("%s", &text)
	for i := 0; i <= len(char); i++ {
		if text == char[i] {
			fmt.Println(getSum(i))
			main()
			break
		}
	}

}
