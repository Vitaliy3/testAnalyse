package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"testSystem/dbStruct"
)

func lp(succesTry int, allTry int) float64 {
	return math.Log(float64((allTry - succesTry) / succesTry))
}
func pBeenrnBaum(lp float64, difficulty float64, c float64, addSens float64) float64 {
	var arg1 float64 = math.Pow(math.E, addSens*(lp-difficulty))
	var arg2 float64 = 1 + math.Pow(math.E, addSens*(lp-difficulty))
	return c + (1-c)*(arg1/arg2)
}

func parsejson(){
	type people struct {
		Graphs           dbStruct.Graphs
		Institutes       dbStruct.Institutes
		Questions        dbStruct.Questions
		Roles            dbStruct.Roles
		Specialities     dbStruct.Specialities
		StudentGroups    dbStruct.StudentGroups
		Subjects         dbStruct.Subjects
		TestResults      dbStruct.TestResults
		TestUsers        dbStruct.TestUsers
		TestUsersSubject dbStruct.TestUsersSubject
		Tests            dbStruct.Tests
		Tickets          dbStruct.Tickets
		UserAnswers      dbStruct.UserAnswers
		Users            dbStruct.Users
	}
	peoples := new(people)
	file, err := os.Open("package.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
	}
	data := make([]byte, fi.Size())
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
		err = json.Unmarshal(data, &peoples)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Success parsing: ", peoples)
	}

}