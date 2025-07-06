package model

import (
    "database/sql"
)

type AssetResponse struct {
    Date         string `json:"date"`
    CurrentValue int    `json:"current_value"`
    CurrentPL    int    `json:"current_pl"`
}

// Step4: 現在の資産評価額、評価損益を計算する関数
func CalculateAsset(db *sql.DB, userID string) (AssetResponse, error) {
    const unitsPerFund = 10000

    // 最新日付（reference_pricesから取得）
    var latestDate string
    err := db.QueryRow(`SELECT MAX(reference_price_date) FROM reference_prices`).Scan(&latestDate)
    if err != nil {
        return AssetResponse{}, err
    }

    // 銘柄ごとの残高（所持口数）を取得
    rows, err := db.Query(`
        SELECT fund_id, SUM(quantity)
        FROM trade_history
        WHERE user_id = ?
        GROUP BY fund_id
    `, userID)
    if err != nil {
        return AssetResponse{}, err
    }
    defer rows.Close()

    totalValue := 0
	// 資産評価額
    for rows.Next() {
        var fundID string
        var qty int
        if err := rows.Scan(&fundID, &qty); err != nil {
            continue
        }

        // 最新の基準価額を取得
        var price int
        err := db.QueryRow(`
            SELECT reference_price
            FROM reference_prices
            WHERE fund_id = ? AND reference_price_date = ?
        `, fundID, latestDate).Scan(&price)
        if err != nil {
            continue // データが無いならスキップ
        }

        totalValue += (price * qty) / unitsPerFund
    }

    // 買付金額の合計を計算
    buyRows, err := db.Query(`
        SELECT fund_id, quantity, trade_date
        FROM trade_history
        WHERE user_id = ?
    `, userID)
    if err != nil {
        return AssetResponse{}, err
    }
    defer buyRows.Close()

    totalCost := 0
    for buyRows.Next() {
        var fundID string
        var qty int
        var tradeDate string
        if err := buyRows.Scan(&fundID, &qty, &tradeDate); err != nil {
            continue
        }

        var price int
        err := db.QueryRow(`
            SELECT reference_price
            FROM reference_prices
            WHERE fund_id = ? AND reference_price_date = ?
        `, fundID, tradeDate).Scan(&price)
        if err != nil {
            continue
        }

        totalCost += (price * qty) / unitsPerFund
    }

    return AssetResponse{
        Date:         latestDate,
        CurrentValue: totalValue,
        CurrentPL:    totalValue - totalCost,
    }, nil
}

// Step5: 指定された日付の資産評価額、評価損益を計算する関数
func CalculateAssetAtDate(db *sql.DB, userID, date string) (AssetResponse, error) {
    const unitsPerFund = 10000

    // 銘柄ごとの残高（その日までの買付を合計）
    rows, err := db.Query(`
        SELECT fund_id, SUM(quantity)
        FROM trade_history
        WHERE user_id = ? AND trade_date <= ?
        GROUP BY fund_id
    `, userID, date)
    if err != nil {
        return AssetResponse{}, err
    }
    defer rows.Close()

    totalValue := 0
	// 資産評価額
    for rows.Next() {
        var fundID string
        var qty int
        if err := rows.Scan(&fundID, &qty); err != nil {
            continue
        }

        var price int
        err := db.QueryRow(`
            SELECT reference_price
            FROM reference_prices
            WHERE fund_id = ? AND reference_price_date = ?
        `, fundID, date).Scan(&price)
        if err != nil {
            continue
        }

        totalValue += (price * qty) / unitsPerFund
    }

    // 買付金額の合計（その日までの買付のみ）
    buyRows, err := db.Query(`
        SELECT fund_id, quantity, trade_date
        FROM trade_history
        WHERE user_id = ? AND trade_date <= ?
    `, userID, date)
    if err != nil {
        return AssetResponse{}, err
    }
    defer buyRows.Close()

    totalCost := 0
    for buyRows.Next() {
        var fundID string
        var qty int
        var tradeDate string
        if err := buyRows.Scan(&fundID, &qty, &tradeDate); err != nil {
            continue
        }

        var price int
        err := db.QueryRow(`
            SELECT reference_price
            FROM reference_prices
            WHERE fund_id = ? AND reference_price_date = ?
        `, fundID, tradeDate).Scan(&price)
        if err != nil {
            continue
        }

        totalCost += (price * qty) / unitsPerFund
    }

    return AssetResponse{
        Date:         date,
        CurrentValue: totalValue,
        CurrentPL:    totalValue - totalCost,
    }, nil
}
