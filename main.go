package main

import (
	"fmt"

	"golang.org/x/net/context"
)

type Messenger interface {
	SendToDevice(ctx context.Context, token string, data map[string]string, notification *Notification) error
	SendToMultipleDevices(ctx context.Context, tokens []string, data map[string]string, notification *Notification) error
	SentToTopic(ctx context.Context, topic string, data map[string]string, notification *Notification) error
}

type Notification struct {
	Title string
	Body  string
}

func main() {
	ctx := context.Background()
	data := map[string]string{
		"subject": "NOTIFICATION",
		"text":    "Hey, ho, let's go",
	}

	testFirebase(ctx, data)
	//testSNS(ctx, data)

	fmt.Print("Done")
}

func testFirebase(ctx context.Context, data map[string]string) {
	firebaseMessenger, err := NewFirebaseMessenger()
	if err != nil {
		panic(err)
	}

	//err = firebaseMessenger.SendToDevice(ctx, "cCeAj2iKRGCIKMdKlYxr9w:APA91bEkQFyM5V3uCh8bAWtDy0xW_XDogdqN7JA1xrzX5tPifW8BZjvowEEmOzKBBguVUEnrTdKLVpfnZTOcPTi8pAqW59HEvvQlOgK_CN6_6LcmmhzEcuX9M9s3mwHgfTqJlnnqrXw_", data, nil)
	err = firebaseMessenger.SentToTopic(ctx, "GATE_210121", data, &Notification{
		Title: "yo",
		Body:  "ho ho",
	})

	if err != nil {
		fmt.Printf("Error sending message to device with firebase: %v", err)
	}
}

func testSNS(ctx context.Context, data map[string]string) {
	snsMessenger, err := NewSNSMessenger()
	if err != nil {
		panic(err)
	}

	err = snsMessenger.SendToDevice(ctx, "arn:aws:sns:us-east-1:343550350117:endpoint/GCM/Presence-Horizon/63ea2f02-3025-324c-8b5d-55f3b17d4d5c", data, nil)
	if err != nil {
		fmt.Printf("Error sending message to device with sns: %v", err)
	}
}
