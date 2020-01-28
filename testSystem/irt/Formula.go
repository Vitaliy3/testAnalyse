package irt

import (
	"database/sql"
	"fmt"
	"gopkg.in/reform.v1"
	"log"
)

var Rdb *reform.DB
var Db *sql.DB

func (v *BinaryMatrix) TcBinaryMatrix() {
	v.Id = make([]int, 11)
	v.Number = make([]int, 9)

	for i := 1; i < 11; i++ {
		v.Id[i] = i
	}
	for i := 1; i < 9; i++ {
		v.Number[i] = i
	}
	v.Value = make([][]int, 11)
	v.Value[0] = []int{0, 0, 1, 0, 1, 0, 1, 1, 1}
	v.Value[1] = []int{0, 0, 0, 0, 1, 0, 0, 1, 1}
	v.Value[2] = []int{0, 0, 0, 0, 0, 1, 0, 1, 0}
	v.Value[3] = []int{0, 0, 0, 0, 1, 0, 1, 1, 1}
	v.Value[4] = []int{1, 0, 0, 0, 1, 0, 0, 1, 0}
	v.Value[5] = []int{1, 1, 1, 0, 0, 0, 1, 1, 1}
	v.Value[6] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	v.Value[7] = []int{1, 0, 1, 1, 0, 1, 1, 1, 1}
	v.Value[8] = []int{0, 0, 0, 1, 0, 0, 1, 1, 0}
	v.Value[9] = []int{1, 0, 1, 0, 1, 1, 1, 1, 1}
	v.Value[10] = []int{1, 1, 1, 0, 1, 1, 1, 1, 1}

}
func (v *BinaryMatrix) Ffull() {
	var counter int
	zadanie := len(v.Number)
	people := len(v.Value)

	for i := 0; i < zadanie; i++ {
		for j := 0; j < people; j++ {
			fmt.Print("-", ":", v.Value[j][i], "	")
			if v.Value[j][i] == 1 {
				counter++
			}
		}
		fmt.Print("  conter:", counter)
		if counter == 11 {
			newSlice := append(v.Value[:i], v.Value[i+1:]...)
			v.Value = newSlice
		}
		counter = 0
		fmt.Println()
	}
}
func (v *BinaryMatrix) CBinaryMatrix() {
	err := Db.QueryRow(`select u.Id,count(u.Id),
case when a.answer = q.true_answer then 1
else 0 end as bValue
from test_users u
left join test_users_subject  tu on u.Id =tu.fk_test_users
left join subjects s on tu.fk_subject = s.Id
left join tests t on s.Id = t.fk_subject
left join tickets tc on t.fk_ticket = tc.Id
left join questions q on q.fk_ticket = tc.Id
left join user_answers a on a.fk_question = q.Id`).Scan(&v.Id, &v.Number, &v.Value)
	if err != nil {
		log.Print(err)
	}
}

func (v *BinaryMatrix) IndividualS() {
	v.IndividualScore = make([]int, 11)
	for i := 0; i < len(v.Id); i++ {
		for j := 0; j < len(v.Number); j++ {
			v.IndividualScore[i] += v.Value[i][j]
		}
	}
}

//func(v *binaryMatrix) cBinaryMatrix()
//{
//}
//func(v *binaryMatrix) cBinaryMatrix()
//{
//}
//func(v *binaryMatrix) cBinaryMatrix()
//{
//}
//func(v *binaryMatrix) cBinaryMatrix()
//{
//}
