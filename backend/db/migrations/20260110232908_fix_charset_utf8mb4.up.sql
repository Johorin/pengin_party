-- データベースの文字セットをutf8mb4に変更
ALTER DATABASE pengin_party CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- usersテーブルの文字セットをutf8mb4に変更
ALTER TABLE users CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- nameカラムの文字セットを明示的にutf8mb4に設定
ALTER TABLE users MODIFY name VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ユーザー名';
