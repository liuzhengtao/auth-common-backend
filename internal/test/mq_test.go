package test

import (
	"context"
	"testing"

	"github.com/liuzhengtao/auth-common-backend/internal/cmd"
)

var (
	ExchangeName       = "direct-selling-admin"
	QueueName          = "direct-selling-queue"
	consumerMessageTip = "监听[直销队列]消费"
	// queueTask          = queue_task.NewQueueTask()
)

type QueueMessageInfo struct {
	MessageId string
	Body      []byte
}

func TestProducer(t *testing.T) {
	// err := queueTask.PutBaseSmsMessage(ctx, ExchangeName, QueueName, &QueueMessageInfo{
	// 	MessageId: "123",
	// 	Body:      []byte("我就随便写点队列内容就行了"),
	// })
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
}

func TestCmd(t *testing.T) {
	cmd.Main.Run(context.Background())
}
