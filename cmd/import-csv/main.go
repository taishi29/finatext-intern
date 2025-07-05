package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/taishi29/finatext-intern/internal/db"
)

func main() {
	// DB接続
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("❌ DB接続失敗:", err)
		return
	}
	defer conn.Close()

	// trade_history の取り込み
	if err := importTradeHistory(conn); err != nil {
		fmt.Println("❌ trade_historyの取り込みに失敗:", err)
		return
	}

	// reference_prices の取り込み
	if err := importReferencePrices(conn); err != nil {
		fmt.Println("❌ reference_pricesの取り込みに失敗:", err)
		return
	}
}

func importTradeHistory(conn *sql.DB) error {
	file, err := os.Open("data/trade_history.csv")
	if err != nil {
		return fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("CSV読み込み失敗: %w", err)
	}

	for i, rec := range records {
		if i == 0 {
			continue // ヘッダーをスキップ
		}

		qty, err := strconv.Atoi(rec[2])
		if err != nil {
			fmt.Println("数量変換エラー:", err)
			continue
		}

		// 重複行はスキップ
		_, err = conn.Exec(`
			INSERT IGNORE INTO trade_history 
			   (user_id, fund_id, quantity, trade_date)
			VALUES (?, ?, ?, ?)
		`, rec[0], rec[1], qty, rec[3])
		if err != nil {
			fmt.Println("INSERTエラー（trade_history）:", err)
			continue
		}

		fmt.Printf("✅ 保存完了（trade_history）: user=%s fund=%s date=%s qty=%d\n",
			rec[0], rec[1], rec[3], qty)
	}

	return nil
}

func importReferencePrices(conn *sql.DB) error {
	file, err := os.Open("data/reference_prices.csv")
	if err != nil {
		return fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("CSV読み込み失敗: %w", err)
	}

	for i, row := range records {
		if i == 0 {
			continue // ヘッダーをスキップ
		}

		price, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("価格変換エラー:", err)
			continue
		}

		// 重複行はスキップ
		_, err = conn.Exec(`
			INSERT IGNORE INTO reference_prices 
			   (fund_id, reference_price_date, reference_price)
			VALUES (?, ?, ?)
		`, row[0], row[2], price)
		if err != nil {
			fmt.Println("INSERTエラー（reference_prices）:", err)
			continue
		}

		fmt.Printf("✅ 保存完了（reference_prices）: fund=%s date=%s price=%d\n",
			row[0], row[2], price)
	}

	return nil
}
