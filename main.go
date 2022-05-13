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

		//get student info
		result, err := models.GetStudentInfo(args)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nNama : %s\nAlamat : %s\nPekerjaan :%s\nAlasan memilih kelas golang:%s\n\n", result.Name, result.Address, result.Work, result.Reason)
		}
	}
}

func LoadDataset() {
	datasets := [][]string{
		{"1", "M. Arsyad R", "Jakarta", "Backend Engineer", "learning new things"},
		{"2", "Tiara D", "Bogor", "Backend Engineer", "learn new programming language"},
		{"3", "Juni Dio K", "Depok", "Backend Engineer", "upgrade skills"},
		{"4", "Tasrifin", "Tangerang", "Backend Engineer", "llearn new programming language"},
		{"5", "Adhitya Febhiakbar", "Jakarta", "Backend Engineer", "upgrade skills"},
		{"6", "Esra Delima M", "Bogor", "Backend Engineer", "upgrade skills"},
		{"7", "M. Avtara", "Depok", "Backend Engineer", "learn new programming language"},
		{"8", "Hamonangan S", "Depok", "Backend Engineer", "upgrade skills"},
		{"9", "Julius M", "Bogor", "Backend Engineer", "learning new things"},
		{"10", "Indra Bayu S", "Tangerang", "Backend Engineer", "learning new things"},
		{"11", "Philip Bryan H", "Jakarta", "Backend Engineer", "learn new programming language"},
		{"12", "Teguh Ainul D", "Jakarta", "Backend Engineer", "upgrade skills"},
		{"13", "Saut Raja M T", "Depok", "Backend Engineer", "learn new programming language"},
		{"14", "Putra Irawan", "Depok", "Backend Engineer", "learning new things"},
		{"15", "Albertus Dwi S", "Bogor", "Backend Engineer", "learning new things"},
		{"16", "Dhany Putra", "Tangerang", "Backend Engineer", "upgrade skills"},
	}

	for _, dataset := range datasets {
		models.AddNewStudent(dataset[0], dataset[1], dataset[2], dataset[3], dataset[4])
	}
}
