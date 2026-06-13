package math

import "fmt"

type Output struct {
	result int
	err    error
}

func CalculateSum(a int, b int) Output {
	sum := a + b
	return Output{
		result: sum,
		err:    nil,
	}
}

type person struct {
	Name string
	Age  *int
	Sex  *string
}

func PersonDetails() *person {
	age := 37
	newAge := &age
	fmt.Println("Age is ", *newAge)
	return &person{Name: "Biswajeet", Age: &age}
}
