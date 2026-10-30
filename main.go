package main

import (
	"fmt"
	"labka8/mathpkg"
)

func main() {
	fmt.Println("Қосу:", mathpkg.Add(7, 5))
	fmt.Println("Азайту:", mathpkg.Subtract(10, 4))
	fmt.Println("Көбейту:", mathpkg.Multiply(6, 6))
	fmt.Println("Бөлу:", mathpkg.Divide(100, 10))
}
