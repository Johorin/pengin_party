FROM golang:1.25.4

WORKDIR /go/pengin_party

# ホットリロードツールのインストール
RUN go install github.com/air-verse/air@latest

# モジュールのインストール＆マイグレーション＆アプリケーション起動
# CMD ["sh", "-c", "go mod tidy && go mod download"]

# CMD ["sh", "-c", "go mod tidy && go mod download && air -c .air.toml"]
# wireのインストール
CMD ["sh", "-c", "go install github.com/google/wire/cmd/wire@latest"]