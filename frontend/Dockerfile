FROM node:latest

WORKDIR /app

# ReactおよびNext.js（App Router対応）をnode:24に適したバージョンでインストール
# RUN npm install --save react@18.3.0 react-dom@18.3.0
# RUN npm install --save next@14.2.3

COPY ./frontend/my-app/package.json /app
COPY ./frontend/my-app/package-lock.json /app

CMD ["sh", "-c", "npm install && npm run dev"]