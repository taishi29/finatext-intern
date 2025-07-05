package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/taishi29/finatext-intern/internal/db"
	"github.com/taishi29/finatext-intern/internal/model"
)

func GetTradeCountHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/") // 先頭末尾の / を除去（例："A1B2C3D4E5/trades"）。
	userID := strings.Split(path, "/")[0] // / で分割し、最初の要素を userID として取得。
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
