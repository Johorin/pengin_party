/**
 * 部屋サービス
 * 部屋関連のビジネスロジックを管理
 */

import { logger } from '../utils/logger.js';

export class RoomService {
  constructor(roomConnectionRepository, messageService) {
    this.roomConnectionRepository = roomConnectionRepository;
    this.messageService = messageService;
  }

  /**
   * クライアントを部屋に参加させる
   * @param {string} roomId - 部屋ID
   * @param {WebSocket} ws - WebSocket接続
   * @param {string|null} previousRoomId - 以前の部屋ID（あれば）
   * @param {object} guest - ゲスト情報（オプション）
   * @returns {string} 新しい部屋ID
   */
  joinRoom(roomId, ws, previousRoomId = null, guest = null) {
    // 以前の部屋から削除
    if (previousRoomId) {
      const previousCount = this.roomConnectionRepository.removeConnection(previousRoomId, ws);
      logger.info(`接続を部屋 ${previousRoomId} から削除。現在の接続数: ${previousCount}`);
    }

    // 新しい部屋に追加
    const newCount = this.roomConnectionRepository.addConnection(roomId, ws);
    logger.info(`接続を部屋 ${roomId} に追加。現在の接続数: ${newCount}`);

    // 部屋の全クライアントに参加者更新を通知
    this.messageService.broadcastToRoom(roomId, {
      type: 'participants-updated',
      roomId: roomId,
      guest: guest
    });

    return roomId;
  }

  /**
   * クライアントを部屋から退出させる
   * @param {string} roomId - 部屋ID
   * @param {WebSocket} ws - WebSocket接続
   */
  leaveRoom(roomId, ws) {
    if (!roomId) {
      return;
    }

    const count = this.roomConnectionRepository.removeConnection(roomId, ws);
    logger.info(`接続を部屋 ${roomId} から削除。現在の接続数: ${count}`);
  }

  /**
   * 部屋の参加者リストを更新して通知
   * @param {string} roomId - 部屋ID
   * @param {Array} participants - 参加者リスト
   */
  notifyParticipantsUpdated(roomId, participants = []) {
    if (!roomId) {
      logger.warn('roomIdが指定されていません');
      return;
    }

    this.messageService.broadcastToRoom(roomId, {
      type: 'participants-updated',
      roomId: roomId,
      participants: participants
    });
  }
}
