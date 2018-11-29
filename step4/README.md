```
[SNS]

マネジメントコンソール → SNS

・トピック
・「新しいトピックの作成」
  ・トピック名： InquiryReceipt
  ・表示名： 省略
  ・「トピックの作成」

・作成したトピックのARNをクリック
・(トピックの詳細)
・「トピックのポリシーの編集」
・「アドバンスド表示」タブ
・以下を追加
---
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": [
        "SNS:Publish"
      ],
      "Resource": "arn:aws:sns:us-east-1:827416367625:InquiryReceipt",
      "Condition": {
        "ArnLike": {
          "AWS:SourceArn": "arn:aws:s3:::inquiry-incoming"
        }
      }
    }
---
・「ポリシーの更新」

・(トピックの詳細)
・「サブスクリプションの作成」
  ・トピックARN： arn:aws:sns:us-east-1:123456789012:InquiryReceipt
  ・プロトコル： AWS Lambda
  ・エンドポイント： arn:aws:lambda:us-east-1:123456789012:function:checkinquity
  ・バージョンまたはエイリアス： default
  ・「サブスクリプションの作成」


[S3のイベント変更]

マネジメントコンソール → S3

・「inquiry-incoming」バケットを選択
・プロパティ
・Events
・既存のイベント(Lambda連係)を選択して「編集」
  ・送信先： Lambda関数 → SNSトピック
  ・SNS： InquiryReceipt
  ・「保存」


[Lambdaの変更]

マネジメントコンソール → Lambda

・「checkinquiry」関数を選択
・トリガーが「SNS」になっていることを確認
  ・トリガー「SNS」をクリック
  ・SNSトピック「InquiryReceipt」が指定されていることを確認

・関数パッケージを更新(アップロード)

・保存

```
