FROM golang:1.25.4

WORKDIR /go/pengin_party

# wireのインストール
RUN go install github.com/google/wire/cmd/wire@v0.7.0

# ホットリロードツールのインストール
RUN go install github.com/air-verse/air@v1.63.4

# CLI migrationツールのインストール @see: https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1

ENV GIN_MODE=debug
ENV APP_ENV=local
ENV TZ="Asia/Tokyo"

# 自動ビルド＆ホットリロード起動
CMD ["air", "-c", ".air.toml"]