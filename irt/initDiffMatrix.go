package irt

import "math"

type InitDifficult struct {
	*BinaryMatrix
	index          []int
	shareRAnswer   []float64
	shareWAnswer   []float64
	lvlDifficult   []float64
	lvlDifficult_S []float64
	middleDiff     []float64
	Dispersia      []float64
}
func (v *InitDifficult) InitDiff() {
	for i, _ := range v.shareWAnswer {
		v.lvlDifficult[i] = math.Log(v.shareWAnswer[i] / v.shareRAnswer[i])

	}
}
func (v *InitDifficult) midPrep() {
	var temp float64
	var numberLen = len(v.Number)
	for i, _ := range v.Number {
		temp = 0
		for _, k := range v.lvlDifficult {
			temp += k
		}
		v.middleDiff[i] = temp / float64(numberLen)
	}
}
func (v *InitDifficult) lvlDiff_S() {
	for i, k := range v.lvlDifficult {
		v.lvlDifficult_S[i] = math.Pow(k, 2)
	}
}
func (v *InitDifficult) Disp() {

}
