/**
 * 部屋接続管理リポジトリ
 * 部屋ごとのWebSocket接続を管理
 */

export class RoomConnectionRepository {
  constructor() {
    // roomId -> Set<WebSocket>
    this.roomConnections = new Map();
  }

  /**
   * 接続を部屋に追加
   * @param {string} roomId - 部屋ID
   * @param {WebSocket} ws - WebSocket接続
   */
  addConnection(roomId, ws) {
    if (!this.roomConnections.has(roomId)) {
      this.roomConnections.set(roomId, new Set());
    }
    this.roomConnections.get(roomId).add(ws);
    return this.roomConnections.get(roomId).size;
  }

  /**
   * 接続を部屋から削除
   * @param {string} roomId - 部屋ID
   * @param {WebSocket} ws - WebSocket接続
   */
  removeConnection(roomId, ws) {
    if (!this.roomConnections.has(roomId)) {
      return 0;
    }

    this.roomConnections.get(roomId).delete(ws);
    
    // 部屋に接続がなくなったら削除
    if (this.roomConnections.get(roomId).size === 0) {
      this.roomConnections.delete(roomId);
      return 0;
    }

    return this.roomConnections.get(roomId).size;
  }

  /**
   * 部屋の接続数を取得
   * @param {string} roomId - 部屋ID
   * @returns {number} 接続数
   */
  getConnectionCount(roomId) {
    return this.roomConnections.get(roomId)?.size || 0;
  }

  /**
   * 部屋の全接続を取得
   * @param {string} roomId - 部屋ID
   * @returns {Set<WebSocket>|undefined} 接続のSet
   */
  getConnections(roomId) {
    return this.roomConnections.get(roomId);
  }

  /**
   * 部屋が存在するか確認
   * @param {string} roomId - 部屋ID
   * @returns {boolean}
   */
  hasRoom(roomId) {
    return this.roomConnections.has(roomId);
  }
}
