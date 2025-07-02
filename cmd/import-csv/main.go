package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/taishi29/finatext-intern/internal/model"
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
	for i, record := range records {
		if i == 0 {
			continue // ヘッダー行をスキップ
		}

		quantity, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Quantityの変換エラー:", err)
			continue
		}

		trade := model.Trade{
			UserID:    record[0],
			FundID:    record[1],
			Quantity:  quantity,
			TradeDate: record[3],
		}

		fmt.Printf("構造体: %+v\n", trade)

	}
}
