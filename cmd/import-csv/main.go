package main

import (
	"fmt"
	"log"

	"github.com/taishi29/finatext-intern/internal/db"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer conn.Close()

	fmt.Println("✅ MySQLに接続成功！")
}
