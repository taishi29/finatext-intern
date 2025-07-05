
# Finatext インターン課題 

# MySQL 操作用コマンド集

このプロジェクトでは、Docker上で立ち上げた MySQL に接続してテーブルの確認やデータの挿入・削除を行います。

---

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
