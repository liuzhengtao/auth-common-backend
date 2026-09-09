package test

import (
	"gitee.com/zhengtao313/lib/queue_task"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

var (
	ExchangeName       = "direct-selling-admin"
	QueueName          = "direct-selling-queue"
	consumerMessageTip = "监听[直销队列]消费"
	queueTask          = queue_task.NewQueueTask()
)

type QueueMessageInfo struct {
	MessageId string
	Body      []byte
}

func TestProducer(t *testing.T) {
	err := queueTask.PutBaseSmsMessage(ctx, ExchangeName, QueueName, &QueueMessageInfo{
		MessageId: "123",
		Body:      []byte("我就随便写点队列内容就行了"),
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestConsumer(t *testing.T) {
	err := queueTask.ConsumerAutoAskHandler(ctx, QueueName, "[direct-test-consumer]", func(msg any) {
		var info *QueueMessageInfo
		err := gconv.Scan(msg, &info)
		if err != nil {
			t.Error(err)
			return
		}
		g.Log().Info(ctx, info.MessageId, string(info.Body))
	})
	if err != nil {
		t.Error(err)
		return
	}
}
