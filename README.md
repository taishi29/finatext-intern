
# Finatext インターン課題（ソフトウェアエンジニアコース）

このリポジトリは、Finatext サマーインターン選考課題（ソフトウェアエンジニアコース）に取り組んだ成果物です。commitメッセージなどは、学んだことを適当に書いてます！

## 使用技術

- Go
- MySQL
- Docker / docker-compose
- chi（軽量ルーティングライブラリ）

 ### 補助ツール？・構成
- Makefile による開発環境統一
- golang-migrate によるマイグレーション管理（migrations/*.sql）

## 🚀 セットアップ手順

## ディレクトリ構成と役割
見よう見まね＆ChatGPTとやり取りして決めたので、適切かどうかはわかりません。。
```
├── cmd/                    # エントリポイント
│   ├── import-csv/         # CSVデータをDBに取り込む用の実行プログラム
│   │   └── main.go
│   └── server/             # Web APIサーバの起動エントリポイント
│       └── main.go
├── data/                   # 課題で使用するCSVデータ
│   ├── reference_prices.csv
│   └── trade_history.csv
├── internal/               # アプリケーションの内部ロジック
│   ├── db/                 # DB接続まわり
│   │   └── conn.go
│   ├── handler/            # HTTPハンドラ（APIのエンドポイント処理）
│   │   ├── asset.go
│   │   └── trade_count.go
│   ├── model/              # DBからのデータ取得や計算ロジック
│   │   ├── asset.go
│   │   ├── prices.go
│   │   └── trade.go
├── migrations/             # DBマイグレーションSQL（テーブル作成など）
│   ├── 0001_create_tables.up.sql
│   └── 0001_create_tables.down.sql
├── docker-compose.yml      # サービス定義（MySQL / アプリ）
├── Dockerfile              # Goアプリのビルド定義
├── go.mod / go.sum         # Goモジュール管理
├── Makefile                # 簡易ビルド・起動・CSVインポート自動化
└── README.md
```

### 1. リポジトリをクローン
```bash
git clone https://github.com/taishi29/finatext-intern.git
cd finatext-intern
```

### 2. コンテナ起動（DB + API）
```bash
make dev/run/import
```
```bash
make dev/run/server
```

## 実装済みAPI一覧

### Step3: トレード数の取得

```
GET /{user_id}/trades
```
- 対象ユーザーのトレード数（件数）を返します

### Step4: 現在の資産評価額と評価損益の取得

```
GET /{user_id}/assets
```
- 最新の基準価額を元に、ユーザーの資産評価額と評価損益を返します

### Step5: 指定日付の資産評価額と評価損益の取得

```
GET /{user_id}/assets?date=YYYY-MM-DD
```
- 指定された日付までの取引と基準価額を元に、当日評価の資産額と損益を返します

##  サンプル実行例

```bash
curl 'http://localhost:8080/F6G7H8I9J0/assets?date=2024-04-01'
```

```json
{
  "date": "2024-04-01",
  "current_value": 13822,
  "current_pl": 44
}
```
---

# MySQL 操作用コマンド集（自分用）

このプロジェクトでは、Docker上で立ち上げた MySQL に接続してテーブルの確認やデータの挿入・削除を行います。

###  1. MySQL に接続する

```bash
docker exec -it finatext-mysql mysql -u user -p
```
- パスワードは `password`

### 2. データベースを選択

```sql
USE finatext;
```

### 3. テーブル一覧を確認

```sql
SHOW TABLES;
```

### 4. テーブルのカラム構成を確認

```sql
DESCRIBE trade_history;
DESCRIBE reference_prices;
```

### 5. テーブルの中身を確認（必要に応じて LIMIT）

```sql
SELECT * FROM trade_history LIMIT 10;
SELECT * FROM reference_prices LIMIT 10;
```

### 6. テーブル内のデータを削除（構造は残す）

```sql
DELETE FROM trade_history;
DELETE FROM reference_prices;
```

> これでテーブルの中の「データだけ」が削除されます。

### 7. 削除確認

```sql
SELECT * FROM trade_history;
SELECT * FROM reference_prices;
```

> `Empty set (0.00 sec)` が表示されれば削除成功！
