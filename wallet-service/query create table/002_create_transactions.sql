CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS transactions (
  id           UUID PRIMARY KEY,
  user_id      BIGINT NOT NULL,
  type         VARCHAR(20) NOT NULL,
  amount       BIGINT NOT NULL,
  status       VARCHAR(20) NOT NULL,
  created_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);