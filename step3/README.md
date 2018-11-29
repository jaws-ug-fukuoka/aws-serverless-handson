```
[SES]

・マネージメントコンソール → SES

Eメールアドレスの検証
・Email Addresses
・Verify a New Email Address
  ・Email Address： 使用するメールアドレス
  ・「Verify This Email Address」
  ・届いたメールのリンクをクリック
  ・Verification Statusが「pending verification」→「verified」になる

Eメール送信テスト
・Email Addresses
・作成したEmail Address Identitiesを選択状態にする
・Send a Test Email


[IAM] (Lambda実行用IAMロール)

・マネージメントコンソール → Lambda
・ロール
・ロールの作成
・1.信頼関係
  ・信頼されたエンティティの種類を選択： AWSサービス
  ・このロールを使用するサービスを選択： Lambda
  ・「次のステップ: アクセス権限」
・2.アクセス権
  ・「Lambda」で検索
  ・ポリシー名：「AWSLambdaBasicExecutionRole」のみにチェック
  ・「次のステップ: タグ」
・3.タグ
  ・「次のステップ: 確認」
・4.確認
  ・ロール名：「Lambda_CheckInquiry」
  ・「ロールの作成」

・作成したポリシーを選択
・アクセス権限
・インラインポリシーの追加
・JSON
---
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::inquiry-incoming/*"
        }
    ]
}
---
・「Review policy」
・名前：S3PutObject
・「Create policy」
・インラインポリシーの追加
・JSON
---
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ses:SendEmail",
            "Resource": "*"
        }
    ]
}
---
・「Review policy」
・名前：SESSendEmail
・「Create policy」
・インラインポリシーの追加
・JSON



[Lambda]

・マネージメントコンソール → Lambda
・関数の作成
・「一から作成」
・名前：「checkinquiry」
・ランタイム： Go 1.x
・ロール： 既存のロールを選択
・既存のロール： Lambda_CheckInquiry
・「関数の作成」

・関数コード
  ・コードエントリタイプ： .zipファイルをアップロード
  ・ランタイム： Go 1.x
  ・ハンドラ： CheckInquiry
  ・関数パッケージ： .zipファイルをアップロードする

・トリガーの追加
  ・S3

・トリガーの設定
  ・バケット： inquiry-incoming
  ・イベントタイプ： オブジェクトの作成(すべて)
  ・プレフィックス： 未設定
  ・サフィックス： 未設定
  ・トリガーの有効化： チェックする
  ・「追加」

・右上の「保存」


[動作確認]

・お問い合わせページから送信する
・S3バケットにJSONファイルが生成されることを確認する
・メールが送信されることを確認する

・CloudWatch → ログ
・ロググループ： /aws/lambda/checkinquiry を選択
・ログストリーム： 最新分を選択
---
08:25:42  START RequestId: 078b4b6f-f21e-11e8-9a1a-2dcbbf9bceab Version: $LATEST
08:25:42  Event receive: Bucket=inquiry-incoming, Key=1543307140345.json
08:25:43  Read from S3 successfully
08:25:43  Send a mail successfully
08:25:43  END RequestId: 078b4b6f-f21e-11e8-9a1a-2dcbbf9bceab
08:25:43  REPORT RequestId: 078b4b6f-f21e-11e8-9a1a-2dcbbf9bceab	Duration: 406.81 ms	Billed Duration: 500 ms Memory Size: 512 MB	Max Memory Used: 33 MB
---
```
