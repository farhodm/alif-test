# EWallet

An API for  test written in Go.

## How to run it:
```shell
git clone git@github.com:farhodm/alif-test.git
cd alif-test
cp .env.example .env
vim .env
# Change the value of APP_ENV from "production" to "local"
# APP_ENV=local
#
# Change the value of variables for DB connection to match yours
#
# Change the value of AUTH_TOKEN_COOKIE_TTL from 100 to whatever minutes you want to stay authorized
# AUTH_TOKEN_COOKIE_TTL=43200
#
# Press Esc button and type this command:
# :x!
sudo -u postgres psql
# - create database test_db;
# - exit
go run ./cmd/console
# Create users and wallets
# Successfully! 
go run ./cmd/api/
# Go to postman
```
