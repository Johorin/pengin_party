/**
 * HTTPコントローラー
 * HTTPエンドポイントのハンドリングを管理
 */

import { logger } from '../utils/logger.js';

export class HttpController {
  constructor(roomService) {
    this.roomService = roomService;
  }

  /**
   * HTTPリクエストを処理
   * @param {http.IncomingMessage} req - HTTPリクエスト
   * @param {http.ServerResponse} res - HTTPレスポンス
   */
  handleRequest(req, res) {
    if (req.method === 'POST' && req.url === '/notify-participant-joined') {
      this.handleNotifyParticipantJoined(req, res);
    } else {
      this.handleNotFound(req, res);
    }
  }

  /**
   * 参加者追加通知エンドポイント
   * @param {http.IncomingMessage} req - HTTPリクエスト
   * @param {http.ServerResponse} res - HTTPレスポンス
   */
  handleNotifyParticipantJoined(req, res) {
    let body = '';

    req.on('data', chunk => {
      body += chunk.toString();
    });

    req.on('end', () => {
      try {
        const data = JSON.parse(body);
        const { roomId, participants } = data;

        if (!roomId) {
          this.sendErrorResponse(res, 400, 'roomId is required');
          return;
        }

        this.roomService.notifyParticipantsUpdated(roomId, participants || []);

        this.sendSuccessResponse(res, { success: true });
      } catch (error) {
        logger.error('JSONパースエラー:', error);
        this.sendErrorResponse(res, 400, 'Invalid JSON');
      }
    });

    req.on('error', (error) => {
      logger.error('リクエストエラー:', error);
      this.sendErrorResponse(res, 500, 'Internal server error');
    });
  }

  /**
   * 404エラーハンドラー
   * @param {http.IncomingMessage} req - HTTPリクエスト
   * @param {http.ServerResponse} res - HTTPレスポンス
   */
  handleNotFound(req, res) {
    this.sendErrorResponse(res, 404, 'Not found');
  }

  /**
   * 成功レスポンスを送信
   * @param {http.ServerResponse} res - HTTPレスポンス
   * @param {object} data - レスポンスデータ
   */
  sendSuccessResponse(res, data) {
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify(data));
  }

  /**
   * エラーレスポンスを送信
   * @param {http.ServerResponse} res - HTTPレスポンス
   * @param {number} statusCode - HTTPステータスコード
   * @param {string} message - エラーメッセージ
   */
  sendErrorResponse(res, statusCode, message) {
    res.writeHead(statusCode, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ error: message }));
  }
}
