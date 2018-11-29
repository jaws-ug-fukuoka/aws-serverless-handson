# aws-serverless-handson

**UNDER CONSTRUCTION**

## 題材

https://speakerdeck.com/morita92hiro/yue-e-10yuan-karazuo-ruserverless-website


## シナリオ(案)

### step-1
「S3で静的Webサイトを作ろう」
* S3(静的Webサイト)

### step-2
「お問い合わせフォームを作ろう」
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM

### step-3
「問い合わせを受けたらメールで通知しよう」
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* Lambda
* SES

### step-4
「SNSを使ってイベント連係を疎結合にしよう」
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* SES

### step-5
「CloudFrontでWebアクセスを高速化しよう」
* CloudFront
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* SES

### step-6
「無料SSL証明書でHTTPS化しよう」
* CloudFront
* ACM
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* SES

### step-7
「WAFで不正アクセスから防御しよう」
* WAF
* CloudFront
* ACM
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* SES

### step-8
「Route 53を使ってダウンしないWebサイトにしよう」
* Route 53
* WAF
* CloudFront
* ACM
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* SES

### step-9
「メール送信をSendGridにパワーアップしよう」
* Route 53
* WAF
* CloudFront
* ACM
* S3(静的Webサイト)
* S3(お問い合わせデータ格納)
* Cognito＋IAM
* SNS
* Lambda
* KMS
* SendGrid
