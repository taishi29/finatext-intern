package handler

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/taishi29/finatext-intern/internal/db"
    "github.com/taishi29/finatext-intern/internal/model"
)

func GetTradeCountHandler(w http.ResponseWriter, r *http.Request) {
    userID := chi.URLParam(r, "user_id")
    if userID == "" {
        http.Error(w, "ユーザーIDが指定されていません。", http.StatusBadRequest)
        return
    }

    conn, err := db.Connect()
    if err != nil {
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    count, err := model.GetTradeCountByUserID(conn, userID)
    if err != nil {
        http.Error(w, "query error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]int{"count": count})
}
