package main

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsMessenger struct {
	*session.Session
}

// TODO: need to setup credentials with IAM https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html
// More examples: https://github.com/aws/aws-sdk-go (including SMS)
func NewSNSMessenger() (Messenger, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		return nil, err
	}
	return snsMessenger{
		sess,
	}, nil
}

func (m snsMessenger) SendToDevice(ctx context.Context, token string, data map[string]string, notification *Notification) error {
	return m.sendMessage(ctx, "", token, data, notification)
}

func (m snsMessenger) SentToTopic(ctx context.Context, topic string, data map[string]string, notification *Notification) error {
	return m.sendMessage(ctx, "", topic, data, notification)
}

func (m snsMessenger) SendToMultipleDevices(ctx context.Context, tokens []string, data map[string]string, notification *Notification) error {
	return errors.New("not implemented")
}

func (m snsMessenger) sendMessage(ctx context.Context, token string, topic string, data map[string]string, notification *Notification) error {
	client := sns.New(m.Session)
	input := &sns.PublishInput{
		Message:          aws.String("Hello world!"), // TODO the format bellow
		MessageStructure: aws.String("json"),         // JSON is just if we want send Firebase stuff on the format bellow, for SMS we need to remove this
	}

	//{
	//	"GCM": "{ \"data\": { \"message\": \"Sample data for Android endpoints\" } }"
	//}

	//{
	//	"GCM": "{ \"notification\": { \"message\": \"Sample notification for Android endpoints\" } }"
	//}

	if token != "" {
		input.SetTargetArn(topic)
	}

	if topic != "" {
		input.SetTopicArn(topic)
	}

	_, err := client.PublishWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
