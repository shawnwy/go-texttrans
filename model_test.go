package texttrans

import (
	"fmt"
	"testing"
)

func TestLoadModel(t *testing.T) {
	matrix, nonPatternProb, err := loadModel("configs/new_model.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(matrix, nonPatternProb)
}

func TestBigramKey(t *testing.T) {
	s := 'S'
	null := '\u0000'
	var x int64
	x = BigramKey(s, null)
	if x != 356482285568 {
		t.Fail()
	}
}

func TestTerminalChar(t *testing.T) {

}
