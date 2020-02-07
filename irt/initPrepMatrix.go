package irt

import "math"

type PIniter interface {
	P() float64
	Disp() float64
}
type InitPreparedness struct {
	index           []int
	IndividualScore []int
	*BinaryMatrix
	shareRAnswer      []float64
	shareWAnswer      []float64
	lvlPreparedress   []float64
	lvlPreparedress_S []float64
	middlePrep        []float64
	Dispersia         []float64
	P                 PIniter
}

func (v *InitPreparedness) InitPrep() {
	for i, _ := range v.shareWAnswer {
		v.lvlPreparedress[i] = math.Log(v.shareRAnswer[i] / v.shareWAnswer[i])
	}
}
func (v *InitPreparedness) midDiff() {
	var temp float64
	var numberLen = len(v.Number)
	for i, _ := range v.Number {
		temp = 0
		for _, k := range v.lvlPreparedress {
			temp += k
		}
		v.middlePrep[i] = temp / float64(numberLen)
	}
}
func (v *InitPreparedness) lvlPrep_S() {
	for i, k := range v.lvlPreparedress {
		v.lvlPreparedress_S[i] = math.Pow(k, 2)
	}
}
func (v *InitPreparedness) Disp() {

}
