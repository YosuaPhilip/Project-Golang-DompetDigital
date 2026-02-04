# Wallet Service (Golang)

Implements:
- Withdraw API
- Balance Inquiry API
- PostgreSQL persistence

## Run (Docker)
```bash
docker compose up --build
```

## Apply migrations
```bash
# open psql inside db container
docker exec -it $(docker ps -qf name=db) psql -U wallet -d walletdb

# then run:
\i /migrations/001_create_wallets.sql
\i /migrations/002_create_transactions.sql
```

Or from host:
```bash
docker exec -i $(docker ps -qf name=db) psql -U wallet -d walletdb < migrations/001_create_wallets.sql
docker exec -i $(docker ps -qf name=db) psql -U wallet -d walletdb < migrations/002_create_transactions.sql
```

## API

### Ensure Wallet (helper)
POST `/api/v1/wallet/ensure`
```json
{ "user_id": 1 }
```

### Balance Inquiry
GET `/api/v1/wallet/balance?user_id=1`

### Withdraw
POST `/api/v1/wallet/withdraw`
```json
{
  "user_id": 1,
  "amount": 50000,
  "reference_id": "wd-20260204-0001"
}
```

Notes:
- Withdraw uses DB transaction + `SELECT ... FOR UPDATE` to avoid race condition.
- Optional idempotency supported using `reference_id`.
