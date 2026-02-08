/**
 * メッセージハンドラー
 * メッセージタイプ別の処理を管理
 */

import { logger } from '../utils/logger.js';

export class MessageHandler {
  constructor(roomService) {
    this.roomService = roomService;
    // 各接続の現在の部屋IDを追跡
    this.connectionRooms = new WeakMap();
  }

  /**
   * メッセージを処理
   * @param {WebSocket} ws - WebSocket接続
   * @param {object} parsedMessage - パース済みメッセージ
   */
  handleMessage(ws, parsedMessage) {
    const { type } = parsedMessage;

    switch (type) {
      case 'join':
        this.handleJoin(ws, parsedMessage);
        break;
      default:
        logger.warn(`未知のメッセージタイプ: ${type}`);
    }
  }

  /**
   * joinメッセージを処理
   * @param {WebSocket} ws - WebSocket接続
   * @param {object} message - メッセージオブジェクト
   */
  handleJoin(ws, message) {
    const { roomId, guest } = message;

    if (!roomId) {
      logger.warn('joinメッセージにroomIdが含まれていません');
      return;
    }

    const previousRoomId = this.connectionRooms.get(ws) || null;
    const newRoomId = this.roomService.joinRoom(roomId, ws, previousRoomId, guest);
    this.connectionRooms.set(ws, newRoomId);
  }

  /**
   * 接続が閉じられた時の処理
   * @param {WebSocket} ws - WebSocket接続
   */
  handleClose(ws) {
    const roomId = this.connectionRooms.get(ws);
    if (roomId) {
      this.roomService.leaveRoom(roomId, ws);
      this.connectionRooms.delete(ws);
    }
  }
}
