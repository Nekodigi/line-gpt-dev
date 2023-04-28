# line-gpt-dev
![LINE GPT](https://user-images.githubusercontent.com/38092328/235027512-e3ce2681-9c70-47ed-ada8-385de1cd93a0.png)

- [Demo](#demo)
- [Blog](#blog)
- [Feature](#feature)
- [TODO](#todo)
- [Setup](#setup)
  - [1. ChatGPT API, Line公式アカウント, Firebaseの準備](#1-chatgpt-api-line公式アカウント-firebaseの準備)
  - [2. 機能の定義](#2-機能の定義)
- [Usage](#usage)
  - [1. 会話モードを使う](#1-会話モードを使う)
  - [2. モード(コマンド)の切り替え](#2-モードコマンドの切り替え)
  - [3. カスタムコマンドの作成](#3-カスタムコマンドの作成)
  - [4. カスタムコマンドを使う](#4-カスタムコマンドを使う)
- [Source Code](#source-code)
  - [Backend](#backend)

## Demo

[https://lin.ee/TRJzQJp](https://lin.ee/TRJzQJp)

## Blog

画像つきの詳細解説はBlogを参照ください。  
[Blogへ移動](https://nekodigi.hatenablog.com/entry/2023/04/28/094650?_ga=2.254117192.1360339689.1682641090-1726183129.1680835625)

## Feature

- 機能を無限に追加できる　
- 会話、翻訳、励まし、要約が初めから使える
- 引数もカスタマイズ可能

## TODO

- カスタム機能入力の簡易化

## Setup

1. ChatGPT API, Line公式アカウント, Firebaseの準備

boilerplateを参考に、Backendのdev.envにsecret等を入力します。

1. 機能の定義

LINEで友達追加すると会話モードに移行しますが、会話モードも定義する必要があります。`go run main.go setup`を実行して会話などの基本的な機能を作成してください。

## Usage

1. 会話モードを使う

友達追加すると自動で会話モードに移行します。何かメッセージを送ると、GhatGPTがそのまま応答します。

1. モード(コマンド)の切り替え

デフォルトでは、[翻訳,励まし,要約]が定義されています。命令名をチャットに入力すると、コマンドが切り替わり、以降のチャットの応答が変化します。

1. カスタムコマンドの作成

カスタムと入力すると、作成が始まります。以下の順番で入力してください。

- 命令名
- 命令文

命令文の%sにはユーザが入力した引数が入ります。最初の%sはユーザーからの入力で置き換えられます。それ以降の%sは、カスタムコマンドに入るときに一度だけ聞かれる引数で置き換えられます。

- 引数

配列で、引数の呼び名を指定します。引数がない場合は空配列[]を指定します。

以下はサンプルコマンドです。

1. カスタムコマンドを使う

コマンド名を入力すると、カスタムコマンドが実行されます。引数がある場合は先に引数を聞かれます。カスタムコマンド実行中は、ユーザーの入力をもとにChatGPTの命令文が自動生成され、それに応じた返答が返ってきます。

## Source Code

- Backend

[https://github.com/Nekodigi/line-gpt-dev](https://github.com/Nekodigi/line-gpt-dev)
