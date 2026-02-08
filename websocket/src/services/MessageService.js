/**
 * メッセージ処理サービス
 * WebSocketメッセージの送信を管理
 */

import { logger } from '../utils/logger.js';

export class MessageService {
  constructor(roomConnectionRepository) {
    this.roomConnectionRepository = roomConnectionRepository;
  }

  /**
   * 部屋の全クライアントにメッセージをブロードキャスト
   * @param {string} roomId - 部屋ID
   * @param {object} message - 送信するメッセージ
   */
  broadcastToRoom(roomId, message) {
    const connections = this.roomConnectionRepository.getConnections(roomId);
    
    if (!connections) {
      logger.warn(`部屋 ${roomId} に接続がありません`);
      return;
    }

    const messageStr = JSON.stringify(message);
    let sentCount = 0;

    connections.forEach((ws) => {
      if (ws.readyState === ws.OPEN) {
        try {
          ws.send(messageStr);
          sentCount++;
        } catch (error) {
          logger.error(`メッセージ送信エラー (roomId: ${roomId}):`, error);
        }
      }
    });

    logger.info(`部屋 ${roomId} にメッセージをブロードキャスト (送信数: ${sentCount}/${connections.size}): ${messageStr}`);
  }

  /**
   * 単一のWebSocket接続にメッセージを送信
   * @param {WebSocket} ws - WebSocket接続
   * @param {object} message - 送信するメッセージ
   */
  sendToConnection(ws, message) {
    if (ws.readyState === ws.OPEN) {
      try {
        ws.send(JSON.stringify(message));
      } catch (error) {
        logger.error('メッセージ送信エラー:', error);
      }
    }
  }
}
