package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type firebaseMessenger struct {
	*firebase.App
}

func NewFirebaseMessenger() (Messenger, error) {
	// Taken from https://console.firebase.google.com/u/0/project/presence-horizon/settings/serviceaccounts/adminsdk
	// DONT GENERATE ANOTHER ONE, ASK AROUND!
	opt := option.WithCredentialsFile("super-secret-firebase-account-key.json")
	config := &firebase.Config{ProjectID: "presence-horizon"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return firebaseMessenger{
		app,
	}, nil
}

// TODO: Firebase errors: https://firebase.google.com/docs/cloud-messaging/send-message#admin

func (m firebaseMessenger) SendToDevice(ctx context.Context, token string, data map[string]string, notification *Notification) error {
	client, err := m.Messaging(ctx)
	if err != nil {
		return err
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data:         data,
		Notification: notification.toFirebaseNotification(),
		Token:        token,
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (m firebaseMessenger) SendToMultipleDevices(ctx context.Context, tokens []string, data map[string]string, notification *Notification) error {
	client, err := m.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.MulticastMessage{
		Data:         data,
		Notification: notification.toFirebaseNotification(),
		Tokens:       tokens,
	}

	_, err = client.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}
	/*
		if br.FailureCount > 0 {
			var failedTokens []string
			for idx, resp := range br.Responses {
					if !resp.Success {
							// The order of responses corresponds to the order of the registration tokens.
							failedTokens = append(failedTokens, registrationTokens[idx])
					}
			}
	*/

	return nil
}

func (m firebaseMessenger) SentToTopic(ctx context.Context, topic string, data map[string]string, notification *Notification) error {
	client, err := m.Messaging(ctx)
	if err != nil {
		return err
	}
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data:         data,
		Notification: notification.toFirebaseNotification(),
		Topic:        topic,
	}

	// Send a message to the devices subscribed to the provided topic.
	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *Notification) toFirebaseNotification() *messaging.Notification {
	if n == nil {
		return nil
	}

	return &messaging.Notification{
		Body:  n.Body,
		Title: n.Title,
	}
}

// TODO: You can group up to 500 messages into a single batch and send them all in a single API call
func (m firebaseMessenger) SendInBatch(ctx context.Context, topic []string, tokens []string, data map[string]string, notification *Notification) error {
	client, err := m.Messaging(ctx)
	if err != nil {
		return err
	}

	// Create a list containing up to 100 messages.
	messages := []*messaging.Message{
		{
			Token: "token",
			Topic: "",
		},
	}

	br, err := client.SendAll(context.Background(), messages)
	if err != nil {
		log.Fatalln(err)
	}

	// See the BatchResponse reference documentation
	// for the contents of response.
	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)

	return nil
}
