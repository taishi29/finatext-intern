package main

import (
	"log"
	"net/http"

	"github.com/taishi29/finatext-intern/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    // Step3: トレード数カウントAPI
    r.Get("/{user_id}/trades", handler.GetTradeCountHandler)

	// Step4＆Step5: 資産評価額と評価損益を返すAPI 
	r.Get("/{user_id}/assets", handler.GetAssetHandler)

    log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}