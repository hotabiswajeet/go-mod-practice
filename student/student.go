package student

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "student.txt"

type Student struct {
	Name  string `json:"name"`
	Roll  int    `json:"roll"`
	Sex   string `json:"sex"`
	Class string `json:"class"`
	Marks string `json:"marks"`
}

func addStudent(name, sex, class, marks string, roll int) *Student {
	return &Student{
		name,
		roll,
		sex,
		class,
		marks,
	}
}

func MaintainStudentData(student []Student) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, stu := range student {
		jsonData, err := json.Marshal(stu)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		jsonData = append(jsonData, '\n')
		if _, err := file.Write(jsonData); err != nil {
			log.Fatal(err)
		}
	}
}

func ReadStudentDataFromFile() []Student {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stud []Student
	for scanner.Scan() {
		line := scanner.Text()
		jsonB := []byte(line)
		var student Student
		err := json.Unmarshal(jsonB, &student)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		stud = append(stud, student)
	}
	return stud
}

func FindTopper(student []Student) Student {
	var topper Student
	marks := 0
	for _, stud := range student {
		mark := extractMarks(stud.Marks)
		if mark > marks {
			marks = mark
			topper = stud
		}
	}
	return topper
}

func extractMarks(markString string) int {
	parts := strings.Split(markString, "/")
	mark, _ := strconv.Atoi(parts[0])
	return mark
}

func updateMarks(mark int) string {
	newMark := strconv.Itoa(mark) + "/100"
	return newMark
}

func StudentAdder() {
	std1 := addStudent("odo", "male", "class II", "20/100", 10)
	std2 := addStudent("litchu", "female", "class II", "90/100", 6)
	std3 := addStudent("globo", "female", "class II", "60/100", 11)
	std4 := addStudent("malli", "female", "class II", "50/100", 21)

	students := []Student{*std1, *std2, *std3, *std4}
	MaintainStudentData(students)
}

func MarkUpdater(students []Student) []Student {
	for i, stud := range students {
		if stud.Name == "odo" {
			newMarks := updateMarks(100)
			students[i].Marks = newMarks
			return students
		}
	}
	return students
}

func FindStudentByRoll(students []Student, roll int) *Student {
	for i, stud := range students {
		if stud.Roll == roll {
			return &students[i]
		}
	}
	return &Student{}
}

func DeleteStudentByRoll(students []Student, roll int) []Student {
	for i, stud := range students {
		if stud.Roll == roll {
			students = append(students[:i], students[i+1:]...)
			break
			// students = students.Delete(students, i, i+1)
		}
	}
	return students
}
