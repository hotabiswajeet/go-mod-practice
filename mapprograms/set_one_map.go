package mapprograms

import "fmt"

type employee struct {
	name    string
	age     string
	details map[string]string
}

func CheckMapFunc() {
	stud := make(map[string]string)
	arr1 := make([]int, 10)
	stud["name"] = "Biswa"
	stud["age"] = "45"
	stud["marks"] = "100"
	arr1[0] = 10
	arr1[1] = 20
	fmt.Println(arr1, stud)
	if _, exists := stud["sex"]; !exists {
		fmt.Println("Key is missing")
	}
	for key, val := range stud {
		fmt.Println(key, val)
	}
}

func SliceStructBehavior() {
	emp := employee{}
	fmt.Println(emp.details == nil)
	var s []int
	s1 := []int{}
	s2 := make([]int, 0)
	fmt.Println(s == nil)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}

func EmptyMapBehavior() {
	var map1 map[int]string
	map2 := make(map[int]string)
	fmt.Println("Just declaration without make", map1 == nil)
	fmt.Println("With make map", map2 == nil)
	empMap := make(map[string]string)
	empMap["dept"] = "Engineering"
	emp1 := employee{name: "Biswa", age: "35", details: empMap}
	fmt.Println("Employee details are", emp1)
}

func StructFilter[T any](input T, filterFunc func(T) bool) T {
	if filterFunc(input) {
		return input
	}
	var zero T
	return zero
}

func MapFilterFunc[T any, K any](input T, key K, filterFunc func(T, K) bool) T {
	if filterFunc(input, key) {
		return input
	}
	var zero T
	return zero
}
