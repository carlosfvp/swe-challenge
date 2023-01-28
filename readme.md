# Project Structure

## `indexer`

Indexer app that push all email files into ZincSearch

---

## `web`

Go web app with Chi that contains:

- API to make searches to ZincSearch
- SPA made with Vue that consumes the API

How to run

1. On the root folder, start ZincSearch with the command `zinc`
1. Go to "indexer" folder and run `go run main.go` to build index the enron mails
1. Go to "web/ui/" folder and build the frontend with `npm install && npm run build`
1. Go to "web" folder and run the API made on go with `go run main.go`

---

## `terraform`

Terraform project to deploy the project in AWS

`In progress`
