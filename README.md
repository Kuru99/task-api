# task-api

Go + PostgreSQL で構築したタスク管理 REST API です．GoやDockerの勉強の為に生成AIを使用して作成しました．

## 概要

タスクの作成・取得・更新・削除（CRUD）を行える REST API です.
Docker でコンテナ化し、Render にデプロイしています

**URL**：https://task-api-g3m6.onrender.com

> ※ 無料プランのため，初回アクセス時に起動まで約 30〜50 秒かかる場合があります．

---

## 使用技術

| 技術 | 用途 |
|---|---|
| Go | サーバーサイド言語 |
| Gin | Web フレームワーク |
| PostgreSQL | データベース |
| Docker | コンテナ化 |
| Render | デプロイ・ホスティング |
| GitHub | ソースコード管理 |

---

## API エンドポイント

| メソッド | パス | 説明 |
|---|---|---|
| GET | /tasks | タスク一覧取得 |
| POST | /tasks | タスク作成 |
| PUT | /tasks/:id | タスク更新 |
| DELETE | /tasks/:id | タスク削除 |

### リクエスト例

**タスク作成（POST /tasks）**
```json
{
  "title": "買い物に行く",
  "done": false
}
```

**レスポンス例**
```json
{
  "id": 1,
  "title": "買い物に行く",
  "done": false
}
```

---

## ローカルでの起動方法

### 必要なもの

- Docker
- Docker Compose


## ディレクトリ構成

```
task-api/
├── main.go              # メインのアプリケーションコード
├── Dockerfile           # Docker イメージの設定
├── docker-compose.yml   # ローカル開発環境の設定
├── init.sql             # DB 初期化 SQL
├── go.mod               # Go モジュール定義
└── go.sum               # 依存関係のチェックサム
```

---

## 工夫した点

- **環境変数で接続情報を管理**：パスワードなどをコードに直書きせず，環境変数から読み込むことでセキュリティに配慮しました．
- **Docker で環境を統一**：どの環境でも同じように動作するよう Docker でコンテナ化しました．
- **マルチステージビルド**：Dockerfile をビルド環境と実行環境に分けることで，最終イメージを軽量化しました．
