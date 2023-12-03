package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

var flagTimer int

func init() {
	flag.IntVar(&flagTimer, "limit", 30, "テストの制限時間(秒単位)")
}

func main() {
	flag.Parse()

	// 問題と回答が記載されたcsvファイルを読み込む
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var answer string
	counter := 0
	passed := 0
	reader := csv.NewReader(file)

	for i := 0; ; i++ {
		record, err := reader.Read()
		// EOFに到達したらループを抜ける
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Q%d: %s = ", i+1, record[0])
		// ターミナルからの入力を受け取る
		fmt.Scan(&answer)

		// 正解数を数える
		if answer == record[1] {
			passed += 1
		}
		counter += 1
	}

	if passed == counter {
		fmt.Printf("あなたは%d問全問正解しました! おめでとうございます🎉", counter)
	} else {
		fmt.Printf("あなたは%d問中%d問正解しました!", counter, passed)
	}
}
