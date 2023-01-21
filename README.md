# Line Bot
* Create a simple Line Bot that can receive and push messages.
* Messages received are stored in MongoDB.

## Setup local development
- [Docker desktop](https://www.docker.com/products/docker-desktop/)
- [Go](https://go.dev/)
- [Homebrew](https://brew.sh/)
- Install ngrok (for macOS users)

   `brew install --cask ngrok`

    用途: 使用ngrok，把外界的請求轉發到你的local

## Setup MongoDB
* Start MongoDB container:
  `make up`

## Got A LINE Bot API devloper account
   Make sure you already registered on [LINE developer console](https://developers.line.biz/console/).
   1. Create new `Messaging Channel`
   2. Get `Channel Secret` on "Basic Setting" tab.
   3. Issue `Channel Access Token` on "Messaging API" tab.
   4. Open LINE Official Account Manager from "Basic Setting" tab.
   5. Go to Reply setting on LINE Official Account Manager, enable "webhook"

## How to run
* Clone this project
  ```
  # Move to your workspace
  cd your-workspace
  
  # Clone this project into your workspace
  git clone git@github.com:karenchuu/go-linebot.git

  # Move to the project root directory
  cd go-linebot
  ```

* Run server:
    ```
    go run main.go
    ngrok http 8080
    ```

* Go to Messaging API on LINE Official Account Manager, input webhookurl. for example: `https://{{YOUR NGROK URL}}/linecallback`

## The Project Layout
```
.
├── config
│   ├── config.yaml
├── db
│   └── mongo.go
├── docker-entrypoint-initdb.d
│   └── mongo-init.js
├── internal
│   ├── api
│   └── models
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── Makefile
```
