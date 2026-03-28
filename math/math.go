package math

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
