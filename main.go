package main

import (
	"assigment1/models"
	"fmt"
	"os"
)

func main() {

	//check if the parameter doesn't exist
	if len(os.Args[1:]) == 0 {
		fmt.Println("[ERROR] parameter not found.")
	} else {
		//get argument
		args := os.Args[1]

		//load data
		LoadDataset()

		students := models.GetAllStudents()
		// fmt.Println(students)

		result, err := models.GetStudentInfo(args, students)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nNama : %s\nAlamat : %s\nPekerjaan :%s\nAlasan memilih kelas golang:%s\n\n", result.Name, result.Address, result.Work, result.Reason)
		}
	}
}

func LoadDataset() {
	datasets := [][]string{
		{"1", "dwi", "jakarta", "dev", "learning 1"},
		{"2", "albert", "bogor", "qa", "learning 2"},
		{"3", "septianto", "depok", "ui", "learning 3"},
	}

	for _, dataset := range datasets {
		models.AddNewStudent(dataset[0], dataset[1], dataset[2], dataset[3], dataset[4])
	}
}
