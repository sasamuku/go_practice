## PostgreSQL インストール
1. `brew install postgresql`
2. `postgres --version`

## DB セットアップ
1. `createuser -P -d gwp4`
2. `createdb gwp4`
3. `psql -U gwp4 -f 13_relational_mappers/setup.sql -d gwp4`
