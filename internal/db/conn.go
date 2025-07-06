package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "user:password@tcp(mysql:3306)/finatext?parseTime=true"

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("DB接続失敗（Open）: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			fmt.Println("DB接続成功")
			return db, nil
		}

		fmt.Printf("DB接続失敗（Ping）: %v\n", err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("DBへの接続リトライ失敗: %w", err)
}
