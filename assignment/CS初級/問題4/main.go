package main

import (
	"fmt"
)

// 学生構造体
type Student struct {
	Name  string
	Score int
}

// 平均点を計算する関数
func averageScore(students []Student) float64 {
	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func main() {
	// 学生データを定義
	students := []Student{
		{"太郎", 70},
		{"花子", 50},
		{"次郎", 80},
	}

	// 平均点計算
	average := averageScore(students)

	// 情報出力
	fmt.Println("学生:", students[0].Name, students[1].Name, students[2].Name)
	fmt.Println("点数:", students[0].Score, students[1].Score, students[2].Score)
	fmt.Println("平均点:", average)

	// 合否判定
	if average >= 60 {
		fmt.Println("判定:合格")
	} else {
		fmt.Println("判定:不合格")
	}
}
