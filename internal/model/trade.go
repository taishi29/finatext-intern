package model

import (
	"database/sql"
)

type Trade struct {
	UserID    string
	FundID    string
	Quantity  int
	TradeDate string // 今は文字列。あとで time.Time にしてもOK
}

func GetTradeCountByUserID(db *sql.DB, userID string) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM trade_history WHERE user_id = ?", userID).Scan(&count)
	return count, err
}
