package main

import (
	"strings"
	"testing"
)

func TestCalc(t *testing.T) {

	tests := []struct {
		data string
		res  int
		res2 int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71, 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51, 51},
		{"2 * 3 + (4 * 5)", 26, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
	}

	for _, v := range tests {
		nospace := strings.ReplaceAll(v.data, " ", "")
		array, _ := translate1(nospace)
		res2 := eval(array)
		if res2 != v.res {
			t.Errorf("Fail (translate1) %s %d should be %d", v.data, res2, v.res)
		}
	}

	for _, v := range tests {
		nospace := strings.ReplaceAll(v.data, " ", "")
		array, _ := translate2(nospace, "")
		res2 := eval(array)
		if res2 != v.res2 {
			t.Errorf("Fail (translate1) %s %d should be %d", v.data, res2, v.res2)
		}
	}

}

func TestEval(t *testing.T) {
	var e = []string{"1", "2", "+", "3", "*", "4", "+", "5", "*", "6", "+"}
	res := eval(e)
	if res != 71 {
		t.Errorf("Fail %d should be %d", res, 71)
	}
}
