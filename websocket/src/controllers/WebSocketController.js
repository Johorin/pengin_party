/**
 * WebSocketコントローラー
 * WebSocket接続のハンドリングを管理
 */

import { logger } from '../utils/logger.js';

export class WebSocketController {
  constructor(messageHandler) {
    this.messageHandler = messageHandler;
  }

  /**
   * WebSocket接続を処理
   * @param {WebSocket} ws - WebSocket接続
   */
  handleConnection(ws) {
    logger.info('新しいWebSocket接続が確立されました');

    ws.on('error', (error) => {
      logger.error('WebSocket error:', error);
    });

    ws.on('message', (data) => {
      try {
        const parsed = JSON.parse(data.toString());
        this.messageHandler.handleMessage(ws, parsed);
      } catch (error) {
        logger.error('メッセージのパースエラー:', error);
      }
    });

    ws.on('close', () => {
      logger.info('WebSocket接続が閉じられました');
      this.messageHandler.handleClose(ws);
    });
  }
}
