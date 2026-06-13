package sliceprograms

import (
	"fmt"
	"runtime"
)

func ReverseSlice(inputSlice []int) {
	length := len(inputSlice)
	j := length - 1
	traverseTill := length/2 - 1
	for i := 0; i <= traverseTill; i++ {
		temp := inputSlice[i]
		inputSlice[i] = inputSlice[j]
		inputSlice[j] = temp
		j--
	}
	fmt.Println("Reversed Slice", inputSlice)
}

func RotateSliceToLeftCheeky(degree int) {
	inputSlice := []int{1, 2, 3, 4, 5}
	length := len(inputSlice)
	fmt.Println(append(inputSlice[degree:length], inputSlice[:degree]...))
}

func reverseSliceLocal(input []int, length int, startIndex int) {
	traverseTill := (length/2 - 1) + startIndex
	j := length - 1 + startIndex
	for i := startIndex; i <= traverseTill; i++ {
		temp := input[i]
		input[i] = input[j]
		input[j] = temp
		j--
	}
}

func RotateSliceToLeft(degree int) {
	inputSlice := []int{1, 2, 3, 4, 5}
	reverseSliceLocal(inputSlice, degree, 0)
	reverseSliceLocal(inputSlice, len(inputSlice)-degree, degree)
	reverseSliceLocal(inputSlice, len(inputSlice), 0)
	fmt.Println("Rotated Slice", inputSlice)
}

func RotateSliceToRight(degree int) {
	inputSlice := []int{1, 2, 3, 4, 5}
	reverseSliceLocal(inputSlice, degree, degree+1)
	reverseSliceLocal(inputSlice, degree+1, 0)
	reverseSliceLocal(inputSlice, len(inputSlice), 0)
	fmt.Println("Rotated Slice", inputSlice)
}

func SliceWithCapacity() {
	slice1 := make([]int, 2, 5)
	var slice2 []int
	//slice2 := []int{}
	fmt.Println("Length of slice is ", len(slice1))
	fmt.Println("Capacity of slice is ", cap(slice1))
	fmt.Println("slice is ", slice1)
	fmt.Println("slice is ", slice2)
}

func SliceSharing() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2 = slice2[:2]
	slice2[1] = 100
	fmt.Println("Slice 1 ", slice1)
	fmt.Println("Slice 2 ", slice2)
}

func RemoveDuplicates() {
	slice1 := []int{10, 40, 30, 20, 15, 10, 20}
	for i := 0; i < len(slice1); i++ {
		for j := i + 1; j < len(slice1); j++ {
			if slice1[i] == slice1[j] {
				slice1 = append(slice1[:j], slice1[j+1:]...)
				j--
			}
		}
	}
	fmt.Println("Removed dupes", slice1)
}

func FilterSlice[T any](inputSlice []T, internal func(T) bool) []T {
	var out []T
	for _, entry := range inputSlice {
		if internal(entry) {
			out = append(out, entry)
		}
	}
	return out
}

func FilterStructs[T any](input T, filterMarks func(T) bool) string {
	if filterMarks(input) {
		return "Yes"
	}
	return "No"
}

func CheckSliceCap() {
	sl := [5]int{}
	var sl1 []int
	fmt.Println("Capacity of slice is ", cap(sl))
	fmt.Println("Slice sl1 is ", sl1 == nil)
	fmt.Println("Slice sl is ", sl)
}

func SpecialArrayCases() {
	s1 := [6]int{1, 2, 3, 4}
	s2 := s1[1:3]
	// Here length is less than capacity of base array
	s3 := append(s2, 99)
	s3[0] = 98
	fmt.Println("Derived array is", s2, cap(s2), &s1[1], &s2[0])
	fmt.Println("New Derived array is", s3, cap(s3))
	fmt.Println("Base array is", s1, cap(s1))
}

func AlwaysUpdateBase() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:3]
	s2[0] = 99
	fmt.Println("New Array is _____", s2, cap(s2), s1)
}

func LengthGreaterThanCap() {
	pc, _, _, _ := runtime.Caller(0)
	s1 := []int{1, 2, 3}
	s2 := s1[1:3]
	s3 := append(s2, 99)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("\n %s: capacity is %d \n", funcName, cap(s2))
	fmt.Printf("\n %s: New aray %v, original array: %v \n", funcName, s3, s1)
}

func CheckWithJustAppend() {
	s1 := make([]int, 3, 6)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	// length of slice is less than cap
	s2 := append(s1, 99)
	s2[0] = 98
	fmt.Println("Length less than cap for append", s1, cap(s1), s2)

	s3 := []int{1, 2, 3}
	s4 := append(s3, 99)
	s4[0] = 98
	fmt.Println("Length equals append", s3, cap(s3), s4, cap(s4))
}
