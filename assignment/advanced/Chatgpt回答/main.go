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

	filenames := []string{"../file1.txt", "../file2.txt", "../file3.txt"}
	var texts []string
	var wg sync.WaitGroup
	ch := make(chan string)

	// goroutineをファイル数分だけ起動
	for _, filename := range filenames {
		wg.Add(1)
		go func(fn string) {
			ch <- readFile(fn)
			wg.Done()
		}(filename)
	}

	// チャンネルから全ファイルのテキストを受信
	for i := 0; i < len(filenames); i++ {
		text := <-ch
		texts = append(texts, text)
	}

	wg.Wait() // 全ての読み込み完了を待機

	// 各ファイルごとに単語数カウントして表示
	for i, text := range texts {
		count := countWords(text)
		fmt.Printf("%s: %d words\n", filenames[i], count)
	}

	// 合計単語数を出力
	totalwordCount := countWords(strings.Join(texts, " "))
	fmt.Println("Total:", totalwordCount, "words")
	fmt.Println("ファイル読み取り処理を終了します")
}

// 共通のファイル読み込み関数
func readFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening", filename)
		return ""
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error reading", filename)
		return ""
	}
	return string(b)
}

// 単語数をカウントする関数
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
