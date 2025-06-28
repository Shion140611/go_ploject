package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("ファイル読み取り処理を開始します")
	var texts []string
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(3)
	go func() {
		ch <- read1()
		wg.Done()
	}()

	go func() {
		ch <- read2()
		wg.Done()
	}()

	go func() {
		ch <- read3()
		wg.Done()
	}()

	// チャネルから3つのテキストを受け取る
	for i := 0; i < 3; i++ {
		text := <-ch
		texts = append(texts, text)
	}
	wordcountfile1 := countWords(read1())
	wordcountfile2 := countWords(read2())
	wordcountfile3 := countWords(read3())

	wg.Wait()
	close(ch)

	totalwordCount := countWords(strings.Join(texts, " "))
	fmt.Println("file1.txt:", wordcountfile1, "words")
	fmt.Println("file2.txt:", wordcountfile2, "words")
	fmt.Println("file3.txt:", wordcountfile3, "words")
	fmt.Println("Total:", totalwordCount, "words")
	fmt.Println("ファイル読み取り処理を終了します")

}

func read1() string {
	// ファイルをOpenする
	f1, err := os.Open("../file1.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f1.Close()
	// 一気に全部読み取り
	b1, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println("read error")
	}
	return string(b1)
}

func read2() string {
	f2, err := os.Open("../file2.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f2.Close()
	b2, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println("read error")
	}
	return string(b2)
}

func read3() string {
	f3, err := os.Open("../file3.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f3.Close()
	b3, err := ioutil.ReadAll(f3)
	if err != nil {
		fmt.Println("read error")
	}
	return string(b3)
}

func countWords(text string) int {
	// 空白文字で文字列を分割
	words := strings.Fields(text)
	// 分割された単語の数を返す
	return len(words)
}
