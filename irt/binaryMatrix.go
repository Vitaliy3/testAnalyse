package irt

import "log"

type BinaryMatrix struct {
	Id              []int
	Number          []int
	Value           [][]int
	IndividualScore [] int
	RightAnswer     []int
	WrongAnswer     []int
}

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
	taskLen := len(v.Number)
	peopleCount := len(v.Value)

	for i := 0; i < taskLen; i++ {
		if v.RightAnswer[i] == taskLen || v.WrongAnswer[i] == 0 {
			for j := 0; j < peopleCount; j++ {
				newSlice := append(v.Value[i][:j], v.Value[i][j+1:]...)
				v.Value[i] = newSlice
				taskLen -= 1
			}
		}
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

func (v *BinaryMatrix) CountRWAnswer() {
	zadanie := len(v.Number)
	people := len(v.Value)
	var counter_r int
	var counter_w int

	for i := 0; i < zadanie; i++ {
		for j := 0; j < people; j++ {
			if v.Value[j][i] == 1 {
				counter_r++
			} else {
				counter_w++
			}
		}

		v.RightAnswer[i] = counter_r
		v.WrongAnswer[i] = counter_w
	}

}
