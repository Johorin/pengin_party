FROM golang:1.25.4

WORKDIR /go/pengin_party

# wireのインストール
RUN go install github.com/google/wire/cmd/wire@latest

# ホットリロードツールのインストール
RUN go install github.com/air-verse/air@latest

# 対話セッションを維持するためにシェルを起動
CMD ["sh"]