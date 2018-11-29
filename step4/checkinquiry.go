package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
)

const emailFrom = "sender@example.com"
const emailTo = "receiver@example.com"

type inquiry struct {
	Date string `json:"Date"`
	Name string `json:"Name"`
	Text string `json:"Text"`
}

type snsmessage struct {
	Records []struct {
		S3 struct {
			Bucket struct {
				Name string `json:"name"`
			} `json:"bucket"`
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func checkinquiry(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		fmt.Printf("Event receive: SNS Message=%s\n", snsRecord.Message)

		var msg snsmessage
		json.Unmarshal(([]byte)(snsRecord.Message), &msg)
		bucketName := msg.Records[0].S3.Bucket.Name
		objectKey := msg.Records[0].S3.Object.Key
		if bucketName == "" || objectKey == "" {
			fmt.Println("Error occurred at reading SNS message")
		}

		inq, err := reads3(bucketName, objectKey)
		if err != nil {
			fmt.Println("Error occurred at S3 GetObject method: ", err)
			return
		}
		fmt.Println("Read from S3 successfully")

		err = sendmail(inq)
		if err != nil {
			fmt.Println("Error occurred at SES SendEmail method: ", err)
			return
		}
		fmt.Println("Send a mail successfully")
	}
}

func reads3(bucket string, key string) (inquiry, error) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		return inquiry{}, err
	}

	defer result.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)

	var inq inquiry
	json.Unmarshal(buf.Bytes(), &inq)

	return inq, nil
}

func sendmail(inq inquiry) error {
	sess := session.Must(session.NewSession())
	svc := ses.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	subject := "お問い合わせを受信しました"
	body := "Webサイトからお問い合わせを受信しました。\n\n" +
		"受付日時: " + inq.Date + "\n" +
		"お名前: " + inq.Name + "\n" +
		"お問い合わせ内容: " + inq.Text

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(emailTo),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(emailFrom),
	}

	_, err := svc.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(checkinquiry)
}
