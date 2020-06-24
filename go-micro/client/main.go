package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"go-learning/go-micro/message"
	"time"
)

func main() {

	service := micro.NewService(
		micro.Name("student_service"),
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
	time.Sleep(50 * time.Second)
}
