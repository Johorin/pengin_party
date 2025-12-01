FROM node:25-alpine

WORKDIR /app

COPY ./websocket/package.json /app
COPY ./websocket/package-lock.json /app

CMD ["sh", "-c", "npm install && node server.js"]