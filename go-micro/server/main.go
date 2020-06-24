package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"go-learning/go-micro/message"
	"log"
)

func main() {
	//创建一个新的服务对象实例
	service := micro.NewService(
		micro.Name("student_service"),
	)

	//服务初始化
	service.Init()

	//注册
	message.RegisterStudentServiceHandler(service.Server(), new(StudentServiceImpl))

	//运行
	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

//学生服务管理实现
type StudentServiceImpl struct {
}

//服务实现
func (ss *StudentServiceImpl) GetStudent(ctx context.Context, request *message.StudentRequest, resp *message.Student) error {
	studentMap := map[string]message.Student{
		"davie":  message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80},
		"steven": message.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   message.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   message.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New(" 请求参数错误,请重新请求。")
	}

	//获取对应的student
	student := studentMap[request.Name]
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*resp = student
		return nil
	}
	return errors.New(" 未查询当相关学生信息 ")
}
