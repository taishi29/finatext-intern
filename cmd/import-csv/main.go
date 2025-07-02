package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// CSVファイルを開く
	file, err := os.Open("data/trade_history.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(file)

	// CSVファイルを読み込む
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// CSVの内容を表示
	for _, record := range records {
		fmt.Println(record)
	}
}
