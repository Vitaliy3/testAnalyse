package irt

type mStruct struct {
	IndividualScore []int
	Id              []int
	Number          []int
	shareRAnswer    []float64
	shareWAnswer    []float64
}
type Disperia interface {
	Disp() float64
}

func P(v mStruct) {
	qCount := len(v.Number)
	for i, _ := range v.Id {
		v.shareRAnswer[i] = float64(v.IndividualScore[i] / qCount)
	}
}

func Q(v mStruct) {
	for i, _ := range v.Id {
		v.shareWAnswer[i] = 1 - v.shareRAnswer[i]
	}
}