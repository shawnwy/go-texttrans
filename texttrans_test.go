package texttrans

import (
	"testing"
)

var (
	tt, _ = NewTextTrans("configs/new_model.json")
)

func TestConstructor(t *testing.T) {
	_, err := NewTextTrans("configs/new_model.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestProbOfEmptyStr(t *testing.T) {
	if tt.Prob("") != 0 {
		t.FailNow()
	}
}

func TestProbOfChineseCharacter(t *testing.T) {
	prob := tt.Prob("文字")
	if prob == 0 {
		t.FailNow()
	}
}

func BenchmarkProb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tt.Prob("one")
	}
}
