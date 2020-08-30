package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var isShowRowNumber = flag.Bool("n", false, "行番号の表示")

func main() {
	flag.Parse()

	// コマンドの引数を取得
	files := flag.Args()

	// コマンドの引数をrangeで回してファイルを開く
	rowNumber := 1
	for _, filepath := range files {
		fp, err := os.Open(filepath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイルを開くことができませんでした:", err)
			return
		}
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			if *isShowRowNumber {
				fmt.Fprintln(os.Stdout, fmt.Sprint(rowNumber, ":"), scanner.Text())
			} else {
				fmt.Fprintln(os.Stdout, scanner.Text())
			}
			rowNumber++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
			return
		}
		fp.Close() // todo: deferを使ってファイルを閉じる

	}
}
