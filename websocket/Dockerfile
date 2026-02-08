FROM node:25-alpine

WORKDIR /app

COPY ./websocket/package.json /app
COPY ./websocket/package-lock.json /app

# 開発環境用: nodemonでホットリロード
CMD ["sh", "-c", "npm install && npm run dev"]