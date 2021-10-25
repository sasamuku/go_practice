## PostgreSQL インストール
1. `brew install postgresql`
2. `postgres --version`

## DB セットアップ
1. `createuser -P -d gwp3`
2. `createdb gwp3`
3. `psql -U gwp3 -f 12_relation/setup.sql -d gwp3`
