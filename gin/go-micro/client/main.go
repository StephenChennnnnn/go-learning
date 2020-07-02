package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/kafka/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-learning/gin/go-micro/message"
	"log"
	"time"
)

func main() {

	service := micro.NewService(
		micro.Name("student_service"),
		micro.Registry(consul.NewRegistry()),
		micro.Broker(kafka.NewBroker()),
	)
	service.Init()

	studentService := message.NewStudentService("student_service", service.Client())
	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println("false")
		fmt.Println(err)
	}
	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)

	brok := service.Options().Broker
	if err := brok.Connect(); err != nil {
		log.Fatal("broker connection failed, error: ", err.Error())
	}
	student := &message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80}
	msgBody, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{
			"name": student.Name,
		},
		Body: msgBody,
	}
	err = brok.Publish("student_service", msg)
	if err != nil {
		log.Fatal("publish failed: %s\n", err.Error())
	} else {
		log.Println("publish succeed.")
	}

	time.Sleep(50 * time.Second)
}
