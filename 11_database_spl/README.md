## PostgreSQL インストール
1. `brew install postgresql`
2. `postgres --version`

## DB セットアップ
1. `createuser -P -d gwp`
2. `createdb gwp`
3. `psql -U gwp -f 11_database/setup.sql -d gwp`
