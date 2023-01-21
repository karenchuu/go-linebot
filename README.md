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
├── main.go
```
## Test 
Test receiving Line Bot messages and pushing messages.

* 傳送任意 Line 訊息，webhook 收到訊息後，會將 user 資訊及訊息分別存入 MongoDB 的 `users` 和 `messages` 資料表

   <img width="469" alt="Screenshot 2023-01-21 at 12 16 18 PM" src="https://user-images.githubusercontent.com/25980598/213842974-f232c0f3-29b7-4524-bf0b-ac7083abb5e5.png">

* 取得 user 資訊 `http://localhost:8080/v1/users`
   <img width="762" alt="Screenshot 2023-01-21 at 12 25 30 PM" src="https://user-images.githubusercontent.com/25980598/213843298-5ed36e57-ddc2-4fe5-8eff-9bfca8966f17.png">

* 取得指定 user 傳送的訊息 `http://localhost:8080/v1/messages?bot_user_id=<bot_user_id>`
   <img width="769" alt="Screenshot 2023-01-21 at 12 26 12 PM" src="https://user-images.githubusercontent.com/25980598/213877304-9255ae02-2567-4f10-bdc9-bfecd00669c0.png">

* 主動推播訊息給 user `http://localhost:8080/v1/sendMessage`
   <img width="758" alt="Screenshot 2023-01-21 at 12 28 19 PM" src="https://user-images.githubusercontent.com/25980598/213843345-3abe0ab6-e645-40bd-bd00-5b6d573bf767.png">
   
   <img width="473" alt="Screenshot 2023-01-21 at 12 29 10 PM" src="https://user-images.githubusercontent.com/25980598/213843370-94c4ce17-d380-4872-9db1-860ca734dc14.png">
   
* 影片版本（測試步驟跟上面一樣）
![demo](https://user-images.githubusercontent.com/25980598/213842809-364e5967-ee7d-4346-8d2d-f7cd3f44aa89.gif)

   
