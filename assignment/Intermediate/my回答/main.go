package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func averageScore(students []Student) float64 {
	var total int
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func main() {
	var student1 = Student{"鈴木", 70}
	var student2 = Student{"田中", 50}
	var student3 = Student{"吉本", 80}
	var students = []Student{student1, student2, student3}

	aberege := averageScore(students)

	fmt.Println("学生:", students[0].Name, students[1].Name, students[2].Name)
	fmt.Println("点数:", students[0].Score, students[1].Score, students[2].Score)
	fmt.Printf("平均点: %.2f\n", aberege)

	if aberege >= 60 {
		fmt.Println("判定:合格")
	} else {
		fmt.Println("判定:不合格")
	}
}
