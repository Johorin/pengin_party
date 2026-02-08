/**
 * WebSocketサーバー エントリーポイント
 * アプリケーションの初期化とサーバー起動を管理
 */

import { WebSocketServer } from 'ws';
import http from 'http';
import { config } from './src/config/server.js';
import { RoomConnectionRepository } from './src/repositories/RoomConnectionRepository.js';
import { MessageService } from './src/services/MessageService.js';
import { RoomService } from './src/services/RoomService.js';
import { MessageHandler } from './src/handlers/MessageHandler.js';
import { WebSocketController } from './src/controllers/WebSocketController.js';
import { HttpController } from './src/controllers/HttpController.js';
import { logger } from './src/utils/logger.js';

// サーバーの作成
const server = http.createServer();
const wss = new WebSocketServer({ server });

// 依存関係の初期化
const roomConnectionRepository = new RoomConnectionRepository();
const messageService = new MessageService(roomConnectionRepository);
const roomService = new RoomService(roomConnectionRepository, messageService);
const messageHandler = new MessageHandler(roomService);
const webSocketController = new WebSocketController(messageHandler);
const httpController = new HttpController(roomService);

// WebSocket接続のハンドリング
wss.on('connection', (ws) => {
  webSocketController.handleConnection(ws);
});

// HTTPリクエストのハンドリング
server.on('request', (req, res) => {
  httpController.handleRequest(req, res);
});

// サーバー起動
server.listen(config.port, () => {
  logger.info(`WebSocket server is running on port ${config.port}`);
});
