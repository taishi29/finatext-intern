package model

type Trade struct {
	UserID    string
	FundID    string
	Quantity  int
	TradeDate string // 今は文字列。あとで time.Time にしてもOK
}
