# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

Go言語で作成されたブロック崩しゲーム。Ebitenライブラリを使用してゲーム開発を行っています。

## 開発コマンド

```bash
# 依存関係の取得
go mod tidy

# ゲームの実行
go run main.go

# ビルド（実行ファイル作成）
go build -o breakout.exe main.go
```

## プロジェクト構成

- `main.go` - メインのゲームロジックとエントリーポイント
- `go.mod` - Go モジュール定義と依存関係

## ゲーム仕様

- 画面サイズ: 800x600
- パドル操作: 左右矢印キー
- ゲーム開始/リスタート: スペースキー
- ブロック配置: 5行12列
- 物理演算: 基本的な反射と衝突判定を実装

## アーキテクチャ

- `Game`構造体: ゲーム状態の管理
- `Paddle`, `Ball`, `Block`構造体: ゲームオブジェクトの定義
- Ebitenゲームインターフェースの実装（Update/Draw/Layoutメソッド）