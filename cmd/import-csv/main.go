package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/taishi29/finatext-intern/internal/db"
	"github.com/taishi29/finatext-intern/internal/model"
)

func main() {
	// DB接続
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("❌ DB接続失敗:", err)
		return
	}
	defer conn.Close()

	// === trade_history.csv を読み込んでDBにINSERT ===
	if err := importTradeHistory(conn); err != nil {
		fmt.Println("❌ trade_historyの取り込みに失敗:", err)
		return
	}

	// === reference_prices.csv を読み込んでDBにINSERT ===
	if err := importReferencePrices(conn); err != nil {
		fmt.Println("❌ reference_pricesの取り込みに失敗:", err)
		return
	}
}

// Trade履歴の取り込み
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

	for i, record := range records {
		if i == 0 {
			continue // ヘッダーをスキップ
		}

		quantity, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("数量変換エラー:", err)
			continue
		}

		trade := model.Trade{
			UserID:    record[0],
			FundID:    record[1],
			Quantity:  quantity,
			TradeDate: record[3],
		}

		_, err = conn.Exec(`
			INSERT INTO trade_history (user_id, fund_id, quantity, trade_date)
			VALUES (?, ?, ?, ?)
		`, trade.UserID, trade.FundID, trade.Quantity, trade.TradeDate)

		if err != nil {
			fmt.Println("INSERTエラー（trade）:", err)
			continue
		}

		fmt.Printf("✅ 保存完了（trade_history）: %+v\n", trade)
	}

	return nil
}

// Reference価格の取り込み
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
			continue
		}

		price, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("価格変換エラー:", err)
			continue
		}

		ref := model.ReferencePrice{
			FundID:             row[0],
			ReferencePriceDate: row[2],
			ReferencePrice:     price,
		}

		_, err = conn.Exec(`
			INSERT INTO reference_prices (fund_id, reference_price_date, reference_price)
			VALUES (?, ?, ?)
		`, ref.FundID, ref.ReferencePriceDate, ref.ReferencePrice)

		if err != nil {
			fmt.Println("INSERTエラー（reference）:", err)
			continue
		}

		fmt.Printf("✅ 保存完了（reference_prices）: %+v\n", ref)
	}

	return nil
}
