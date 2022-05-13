package models

import (
	"fmt"
	"strconv"
)

type Student struct {
	Id      string
	Name    string
	Address string
	Work    string
	Reason  string
}

var students = []Student{}

func GetStudentInfo(id string, students []Student) (Student, error) {
	idConverted, _ := strconv.Atoi(id)
	if idConverted <= 0 {
		return Student{}, fmt.Errorf("[ERROR] id = %d is not found. please use another id", idConverted)
	}
	return students[idConverted-1], nil
}

func AddNewStudent(id string, name string, address string, work string, reason string) {
	students = append(students, Student{
		Id:      id,
		Name:    name,
		Address: address,
		Work:    work,
		Reason:  reason,
	})
}

func GetAllStudents() []Student {
	return students
}
