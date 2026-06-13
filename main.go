package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/hotabiswajeet/go-mod-practice/filterexercise"
	"github.com/hotabiswajeet/go-mod-practice/mapprograms"
	"github.com/hotabiswajeet/go-mod-practice/math"
	"github.com/hotabiswajeet/go-mod-practice/sliceprograms"
	"github.com/hotabiswajeet/go-mod-practice/student"
)

type classStudent struct {
	name  string
	age   int
	marks int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go app")
}

func main() {

	mapFilterGenImpl()
	mapFilterByValImpl()
	mapFilterByStructImpl()
	mapFilterByMapImpl()
	mapFilterAnyImpl()

	http.HandleFunc("/", handler)

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}

func archiveFromMain() {

	person := math.PersonDetails()
	fmt.Printf("The name of the person is %s and age is %d\n", person.Name, *person.Age)
	// Add students
	student.StudentAdder()

	// Read student data
	studs := student.ReadStudentDataFromFile()

	// update marks
	studs = student.MarkUpdater(studs)

	// find topper
	topper := student.FindTopper(studs)

	stud := student.FindStudentByRoll(studs, 10)

	stud.Marks = "60/100"

	studs = student.DeleteStudentByRoll(studs, 11)

	// insert updated data
	student.MaintainStudentData(studs)

	sliceprograms.ReverseSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	sliceprograms.RotateSliceToLeftCheeky(2)

	sliceprograms.RotateSliceToLeft(2)

	sliceprograms.RotateSliceToRight(2)

	sliceprograms.SliceWithCapacity()

	sliceprograms.SliceSharing()

	sliceprograms.RemoveDuplicates()

	sliceprograms.CheckSliceCap()

	sliceprograms.SpecialArrayCases()

	sliceprograms.AlwaysUpdateBase()

	sliceprograms.LengthGreaterThanCap()

	sliceprograms.CheckWithJustAppend()

	output := sliceprograms.FilterSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) bool {
		return v%2 == 0
	})

	fmt.Println("Even numbers", output)

	newStudent := &classStudent{name: "Biswa", age: 37, marks: 90}

	stOut := sliceprograms.FilterStructs(*newStudent, func(st classStudent) bool {
		return st.marks >= 90
	})

	fmt.Printf("\nIs %s a good student? Ans: %s\n", newStudent.name, stOut)

	fmt.Printf("%+v\n", student.FindStudentByRoll(studs, 10))

	fmt.Printf("\nTopper is %s and with marks %s\n", topper.Name, topper.Marks)

	mapprograms.CheckMapFunc()

	mapprograms.SliceStructBehavior()

	mapprograms.EmptyMapBehavior()

	structFilterImpl()

	mapFilterImpl()

}

func mapFilterGenImpl() {
	employees := map[string]int{
		"John":  30,
		"Biswa": 35,
		"Alice": 25,
	}
	output := filterexercise.MapFilterByKey(employees, func(value int) bool {
		return value > 30
	})

	fmt.Println("Map Filter output", output)
}

func structFilterImpl() {
	person := struct {
		name string
		age  int
	}{name: "Biswajeet", age: 37}

	output := mapprograms.StructFilter(person, func(p struct {
		name string
		age  int
	}) bool {
		return p.name == "Biswa"
	})

	fmt.Println("Output of struct filter", output)

}

func mapFilterImpl() {
	map1 := make(map[string]any)
	map1["Emp1"] = struct {
		name string
		age  int
	}{name: "Anirban", age: 37}
	map1["Emp2"] = []string{"Biswa", "35"}
	key := "Emp2"
	output := mapprograms.MapFilterFunc(map1, key, func(m map[string]any, k string) bool {
		field := m[k]
		val := reflect.ValueOf(field)
		if val.Kind() == reflect.Struct {
			name := val.FieldByName("name")
			if name.IsValid() {
				return name.String() == "Anirban"
			}
		}
		return false
	})

	fmt.Println("Output of map filter", output)
}

func mapFilterByValImpl() {
	data := map[string]any{
		"Emp1": "Biswa",
		"Emp2": 35,
		"Emp3": true,
		"Emp4": "Anirban",
	}
	output := filterexercise.MapFilterByVal(data, func(val any) bool {
		value := reflect.ValueOf(val)
		return value.Kind() == reflect.String
	})

	fmt.Println("The Map with string values", output)
}

func mapFilterByStructImpl() {
	employees := []any{
		struct {
			name string
			age  int
		}{
			name: "Biswa",
			age:  37,
		},
		struct {
			name string
			age  int
		}{
			name: "Anirban",
			age:  38,
		},
	}

	name := "Rohit"

	exists := filterexercise.MapFilterByStruct(employees, name, func(emp any, n string) bool {
		val := reflect.ValueOf(emp)
		if val.Kind() == reflect.Struct {
			nameField := val.FieldByName("name")
			if nameField.IsValid() {
				return nameField.String() == n
			}
		}
		return false
	})
	fmt.Printf("\n %s name exists %v\n", name, exists)
}

func mapFilterByMapImpl() {
	employee := map[string]any{
		"name":   "Biswa",
		"age":    37,
		"emp_id": 101,
	}
	targetName := "Biswa"
	exists := filterexercise.MapFilterByMap(employee, func(emp map[string]any) bool {
		return emp["name"] == targetName
	})

	fmt.Printf("\n %s name exists %v\n", targetName, exists)
}

func mapFilterAnyImpl() {
	type employee struct {
		name string
		age  int
	}
	inputMap := map[string]any{
		"Emp1": employee{"Rubab", 37},
		"Emp2": employee{"Adeel", 40},
		"Emp3": employee{"Musrafa", 29},
		"Emp4": employee{"Sharjeena", 25},
	}

	output := filterexercise.MapFilterAny(inputMap, func(val any) bool {
		valField := reflect.ValueOf(val)
		if valField.Kind() == reflect.Struct {
			field := valField.FieldByName("age")
			if field.IsValid() {
				if field.Kind() == reflect.Int {
					return field.Int() > 35
				}
			}
		}
		return false
	})

	fmt.Println("The output map is ", output)
}
