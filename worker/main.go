package main

import (
	"broker_queues/common"
	"broker_queues/generated/message"
	"log"

	"context"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: common.BrokerAddress,
	})
	ctx := context.Background()

	msg := &message.Message{Data: "empty"}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	err = rdb.Publish(ctx, common.BrokerMainChannel, data).Err()
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}
