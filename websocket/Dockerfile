FROM node:latest

WORKDIR /app

COPY ./websocket/package.json /app
COPY ./websocket/package-lock.json /app

RUN npm install

CMD ["node", "server.js"]