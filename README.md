# EWallet

An API for  test written in Go.

## How to run it:
```shell
git clone git@github.com:farhodm/alif-test.git
cd alif-test
cp .env.example .env
# Change the value of variables for DB connection to match yours
# Press Esc button and type this command:
# :x!
sudo -u postgres psql
# - create database test_db;
# - exit
go run ./cmd/console
# Wait until seeding will finish
# Message "Successfully!" text is okay.
go run ./cmd/api/
# Go to postman
```
