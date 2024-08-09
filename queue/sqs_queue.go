package queue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Queue interface {
	SendMessage(ctx context.Context, message string) error
	ReceiveMessages(ctx context.Context) ([]types.Message, error)
}

type sqsQueue struct {
	client *sqs.Client
	url    string
}

func NewSQSQueue(client *sqs.Client, url string) Queue {
	return &sqsQueue{client: client, url: url}
}

func (q *sqsQueue) SendMessage(ctx context.Context, message string) error {
	_, err := q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(q.url),
		MessageBody: aws.String(message),
	})
	return err
}

func (q *sqsQueue) ReceiveMessages(ctx context.Context) ([]types.Message, error) {
	result, err := q.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(q.url),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     20,
	})
	if err != nil {
		return nil, err
	}
	return result.Messages, nil
}
