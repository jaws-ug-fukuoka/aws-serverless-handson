```
[S3バケット作成] (お問い合わせデータ格納用)

・マネジメントコンソール → S3

・[バケットを作成する]
  ・1 名前とリージョン
    ・バケット名： 一意な名前 (例：「inquiry-incoming」)
    ・リージョン： 任意 (例：米国東部(バージニア北部))
    ・既存のバケットから設定をコピー： (未入力)
    ・[次へ]
  ・2 オプションの設定
    ・[次へ]
  ・3 アクセス許可の設定
    ・[次へ]
  ・4 確認
    ・[バケットを作成]

・作成したバケットを選択
  ・アクセス権限
  ・CORSの設定
---
<?xml version="1.0" encoding="UTF-8"?>
<CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
    <CORSRule>
        <AllowedOrigin>http://aws-serverless-handson.s3-website-us-east-1.amazonaws.com</AllowedOrigin>
        <AllowedMethod>GET</AllowedMethod>
        <AllowedMethod>PUT</AllowedMethod>
        <AllowedHeader>*</AllowedHeader>
    </CORSRule>
</CORSConfiguration>
---


[Cognito]

・マネジメントコンソール → Cognito
・[IDプールの管理]
・[新しいIDプールの作成]
・ステップ1: IDプールを作成する
  ・IDプール名： 任意 (例：「InquiryIdPool」)
  ・認証されていないID
    ・認証されていないIDに対してアクセスを有効にする： チェックを入れる
  ・認証プロバイダー
    ・(未設定)
  ・[プールの作成]
・ステップ2: アクセス権限の設定
  ・[許可]

・サンプルコード
  ・プラットフォーム： JavaScript
  ・AWS認証情報の取得
    ・「IdentityPoolId」を書き控える
      us-east-1:d0181d3b-fc20-4e6d-a9a3-8ea0a9348455


[IAM]

・ロール
・ロール名「Cognito_＜IDプール名＞Unauth_Role」を選択
・アクセス権限
・インラインポリシーの追加
・JSON
---
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::inquiry-incoming/*"
        }
    ]
}
---
・「Review policy」
・名前：S3GetObject
・「Create policy」


[S3] (静的Webサイト用)

・「index.html」の編集
  ---
  const IdentityPoolId = 'us-east-1:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx';
  const BucketRegion = 'us-east-1';
  const BucketName = 'inquiry-incoming';
  ---
・ルートに「index.html」を置く


[動作確認]

・「エンドポイント」欄を確認する
  http://aws-serverless-handson.s3-website-us-east-1.amazonaws.com

・WebブラウザでエンドポイントのURLへアクセス

・何か入力して「送信」

・「お問い合わせを承りました」と表示されることを確認

・お問い合わせデータ格納用S3バケットにJSONファイルが格納されていることを確認
```
