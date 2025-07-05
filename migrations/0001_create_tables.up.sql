CREATE TABLE IF NOT EXISTS trade_history (
  user_id    VARCHAR(10) NOT NULL,
  fund_id    VARCHAR(6)  NOT NULL,
  quantity   INT         NOT NULL,
  trade_date DATE        NOT NULL,
  PRIMARY KEY (user_id, fund_id, trade_date)
);

CREATE TABLE IF NOT EXISTS reference_prices (
  fund_id              VARCHAR(6)  NOT NULL,
  reference_price_date DATE        NOT NULL,
  reference_price      INT         NOT NULL,
  PRIMARY KEY (fund_id, reference_price_date)
);
