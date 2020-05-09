# Firebase

## Setup

To create a firebase project and add it to an android device: https://firebase.google.com/docs/android/setup?authuser=0

For specifics on Cloud messaging: https://firebase.google.com/docs/cloud-messaging/android/client?authuser=0

Code source as example: https://github.com/firebase/quickstart-android/tree/master/messaging

Tutorial with code-source and explanation: https://codelabs.developers.google.com/codelabs/advanced-android-kotlin-training-notifications-fcm/index.html?index=..%2F..index#0

The server side of things (including go which we used here) is https://firebase.google.com/docs/cloud-messaging/send-message?authuser=0#send-messages-to-specific-devices   

# Amazon SNS

## Setup

To start the setup: https://aws.amazon.com/premiumsupport/knowledge-center/create-android-push-messaging-sns/

To add a device to amazon, need to run the device with the firebase setup, and get the token from `onNewToken` method... them https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html

Having an endpoint a manual test can be done to validate

This is the Go example of how handle stuff, but there's a JAVA SDK too: https://docs.aws.amazon.com/sns/latest/dg/mobile-platform-endpoint.html

For the go stuff: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-sns-with-go-sdk.html