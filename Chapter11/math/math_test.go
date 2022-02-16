package math

import "testing"

type testPair struct {
	values  []float64
	average float64
}

var pairs = []testPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	//v := 0.0
	for _, pair := range pairs {
		value := Average(pair.values)
		if value != pair.average {
			t.Error("For:", pair.values,
				"expected:", pair.average,
				"got", value)
		}
	}
}


//func TestAverage(t *testing.T) {
//	v := 0.0
//	v = Average([]float64{1, 2})
//	if v != 1.5 {
//		t.Error("Expected 1.5, got", v)
//	}
//}
