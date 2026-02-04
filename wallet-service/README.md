# Project-Golang-DompetDigital
# Wallet Service (Golang)

Implements:
- Withdraw API
- Balance Inquiry API
- PostgreSQL persistence

## Run (Docker)
```bash
docker compose up --build
001_create_wallets
```

## query create table
```run query for create table in database
001_create_wallets
002_create_transactions
```

## API

### Balance Inquiry
GET `/api/v1/wallet/balance?user_id=1`

### Withdraw
POST `/api/v1/wallet/withdraw`
```json
{
  "user_id": 1,
  "amount": 50000,
}
```

Notes:
- Withdraw uses DB transaction + `SELECT ... FOR UPDATE` to avoid race condition.
