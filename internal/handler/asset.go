package handler

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/taishi29/finatext-intern/internal/db"
    "github.com/taishi29/finatext-intern/internal/model"
)

// Step4: 現在の資産評価額、評価損益を返す関数
func GetAssetHandler(w http.ResponseWriter, r *http.Request) {
    userID := chi.URLParam(r, "user_id")
    date := r.URL.Query().Get("date")

    if date != "" {
        GetAssetAtDateHandler(w, r)
        return
    }

	if userID == "" {
        http.Error(w, "ユーザーIDが指定されていません。", http.StatusBadRequest)
        return
    }
    conn, err := db.Connect()
    if err != nil {
        http.Error(w, "データベースエラー", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    asset, err := model.CalculateAsset(conn, userID)
    if err != nil {
        http.Error(w, "計算失敗", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(asset)
}

// Step5: 指定された日付の資産評価額、評価損益を返す関数
func GetAssetAtDateHandler(w http.ResponseWriter, r *http.Request) {
    userID := chi.URLParam(r, "user_id")
    date := r.URL.Query().Get("date") // クエリパラメータで指定された日付

	if userID == "" {
        http.Error(w, "ユーザーIDが指定されていません。", http.StatusBadRequest)
        return
    }
    if date == "" {
        http.Error(w, "日付が指定されていません。?date=YYYY-MM-DD の形式で指定してください。", http.StatusBadRequest)
        return
    }

    conn, err := db.Connect()
    if err != nil {
        http.Error(w, "データベースエラー", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    asset, err := model.CalculateAssetAtDate(conn, userID, date)
    if err != nil {
        http.Error(w, "計算失敗: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(asset)
}