package mq

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/mq/amqpclt"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

var AddActor *amqpclt.Actor
var DelActor *amqpclt.Actor

// InitMq to init rabbitMQ
func InitMq() {
	amqpConn, err := amqp.Dial(fmt.Sprintf(shared.RabbitMqURI, shared.MQUser, shared.MQPassword, shared.MQHost, shared.MQPort))
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	AddActor, err = amqpclt.NewActor(amqpConn, "follow_add")
	if err != nil {
		klog.Fatal("cannot create add actor", err)
	}
	//开启消费监听
	go AddActor.Consumer(context.Background())
	DelActor, err = amqpclt.NewActor(amqpConn, "follow_del")
	if err != nil {
		klog.Fatal("cannot create del actor", err)
	}
	//开启消费监听
	go DelActor.Consumer(context.Background())
}
