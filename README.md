# 🧱 Go Breakout Game

Go言語とEbitenライブラリで作成したブロック崩しゲーム。WebAssemblyでブラウザ上でも動作します。

## 🎮 プレイ方法

### オンラインでプレイ
**👉 [こちらでプレイできます](https://maniax-jp.github.io/go-breakout-game/)**

### ローカルで実行
```bash
# リポジトリをクローン
git clone https://github.com/maniax-jp/go-breakout-game.git
cd go-breakout-game

# 依存関係を取得
go mod tidy

# ゲームを実行
go run main.go
```

## 🕹️ 操作方法

- **スペースキー**: ゲーム開始/リスタート
- **左右矢印キー**: パドル操作

## 🎯 ゲームの目的

- パドルでボールを跳ね返して、すべてのブロックを破壊する
- ボールを下に落とさないよう注意！
- すべてのブロックを破壊すればクリア

## 🛠️ 技術スタック

- **言語**: Go 1.21+
- **ゲームエンジン**: [Ebiten](https://ebitengine.org/)
- **デプロイ**: GitHub Actions + GitHub Pages
- **WebAssembly**: ブラウザ対応

## 🏗️ ビルド

### デスクトップ版
```bash
go build -o breakout.exe main.go
```

### WebAssembly版
```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```

## 🚀 自動デプロイ

GitHub Actionsを使用して、mainブランチへのプッシュ時に自動的にWASMビルドとGitHub Pagesへのデプロイが実行されます。

## 📝 ライセンス

MIT License

## 🤝 コントリビューション

プルリクエストやイシューの報告を歓迎します！