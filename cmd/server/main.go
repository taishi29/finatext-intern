package main

import (
	"log"
	"net/http"

	"github.com/taishi29/finatext-intern/internal/handler"
)

func main() {
	// ルーティング設定（URLと関数の対応づけ）
	http.HandleFunc("/", handler.GetTradeCountHandler)

	// ポート8080でHTTPサーバーを起動して、もし失敗したら（Listenできなかったら）エラーを出して終了する
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
