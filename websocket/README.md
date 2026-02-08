# WebSocket Server

このプロジェクトは、リアルタイム通信を提供するWebSocketサーバーです。部屋（Room）ごとの接続管理とメッセージブロードキャスト機能を提供します。

## アーキテクチャ

このプロジェクトは、保守性と拡張性を重視したレイヤードアーキテクチャを採用しています。

### ディレクトリ構造

```
websocket/
├── src/
│   ├── config/              # 設定ファイル
│   │   └── server.js        # サーバー設定（ポート、環境変数など）
│   │
│   ├── repositories/        # データアクセス層
│   │   └── RoomConnectionRepository.js  # 部屋接続管理のリポジトリ
│   │
│   ├── services/            # ビジネスロジック層
│   │   ├── RoomService.js   # 部屋関連のビジネスロジック
│   │   └── MessageService.js  # メッセージ処理サービス
│   │
│   ├── controllers/         # プレゼンテーション層
│   │   ├── WebSocketController.js  # WebSocket接続ハンドリング
│   │   └── HttpController.js  # HTTPエンドポイント
│   │
│   ├── handlers/            # メッセージハンドラー
│   │   └── MessageHandler.js  # メッセージタイプ別の処理
│   │
│   ├── utils/               # ユーティリティ
│   │   └── logger.js        # ロギングユーティリティ
│   │
│   └── server.js            # エントリーポイント（最小限のコード）
│
├── package.json
├── Dockerfile
└── README.md
```

### レイヤーの説明

#### 1. Config Layer (`src/config/`)
- **責務**: アプリケーションの設定を管理
- **ファイル**: `server.js` - ポート番号、環境変数などの設定

#### 2. Repository Layer (`src/repositories/`)
- **責務**: データの永続化とアクセスを管理
- **ファイル**: `RoomConnectionRepository.js` - 部屋ごとのWebSocket接続を管理するMapの操作を提供

#### 3. Service Layer (`src/services/`)
- **責務**: ビジネスロジックを実装
- **ファイル**:
  - `RoomService.js` - 部屋への参加・退出、参加者通知などのビジネスロジック
  - `MessageService.js` - メッセージの送信、ブロードキャスト処理

#### 4. Controller Layer (`src/controllers/`)
- **責務**: 外部からのリクエストを受け取り、適切なサービスに委譲
- **ファイル**:
  - `WebSocketController.js` - WebSocket接続の確立、メッセージ受信、接続切断の処理
  - `HttpController.js` - HTTPエンドポイント（`/notify-participant-joined`）の処理

#### 5. Handler Layer (`src/handlers/`)
- **責務**: メッセージタイプに応じた処理を分岐
- **ファイル**: `MessageHandler.js` - `join`などのメッセージタイプ別の処理を実装

#### 6. Utils Layer (`src/utils/`)
- **責務**: 共通ユーティリティ関数を提供
- **ファイル**: `logger.js` - 統一されたログ出力機能

### データフロー

```
1. WebSocket接続確立
   ↓
   WebSocketController.handleConnection()
   ↓
2. メッセージ受信
   ↓
   MessageHandler.handleMessage()
   ↓
3. ビジネスロジック実行
   ↓
   RoomService.joinRoom()
   ↓
4. データアクセス
   ↓
   RoomConnectionRepository.addConnection()
   ↓
5. メッセージ送信
   ↓
   MessageService.broadcastToRoom()
```

### 依存関係の方向

```
server.js (エントリーポイント)
  ↓
Controllers → Handlers → Services → Repositories
  ↓
Utils (共通ユーティリティ)
```

- **上位レイヤーは下位レイヤーに依存**する一方で、**下位レイヤーは上位レイヤーに依存しない**
- これにより、各レイヤーの独立性が保たれ、テストや変更が容易になります

## 主な機能

### WebSocket機能

- **部屋への参加**: クライアントが`join`メッセージを送信すると、指定された部屋に参加します
- **自動退出**: 接続が切断されると、自動的に部屋から退出します
- **ブロードキャスト**: 部屋の全クライアントにメッセージを送信します

### HTTP API

- **POST `/notify-participant-joined`**: 部屋の参加者リストが更新された際に、WebSocketクライアントに通知します

## 使用方法

### 開発環境での起動

```bash
npm run dev
```

### 本番環境での起動

```bash
npm start
```

### 環境変数

- `PORT`: サーバーのポート番号（デフォルト: 3001）
- `NODE_ENV`: 実行環境（`development` / `production`）

## 拡張方法

### 新しいメッセージタイプを追加する場合

1. `src/handlers/MessageHandler.js`の`handleMessage()`メソッドに新しいケースを追加
2. 必要に応じて新しいハンドラーメソッドを実装
3. ビジネスロジックが必要な場合は`src/services/`にサービスを追加

### 新しいHTTPエンドポイントを追加する場合

1. `src/controllers/HttpController.js`の`handleRequest()`メソッドに新しいルートを追加
2. 新しいハンドラーメソッドを実装

### 新しいサービスを追加する場合

1. `src/services/`に新しいサービスファイルを作成
2. 必要に応じて`server.js`でインスタンス化し、依存関係を注入

## テスト

各レイヤーが独立しているため、単体テストが容易です：

- Repository層: データアクセスのテスト
- Service層: ビジネスロジックのテスト
- Controller層: リクエスト/レスポンスのテスト

## 利点

1. **保守性**: 各モジュールが単一責任を持ち、変更箇所が明確
2. **拡張性**: 新機能の追加が容易（適切なレイヤーに追加するだけ）
3. **テスト容易性**: 各レイヤーを独立してテスト可能
4. **可読性**: コードの構造が明確で、理解しやすい
5. **再利用性**: 各サービスやユーティリティを他の場所でも再利用可能
