package texttrans

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	tt, err := NewTextTrans("configs/new_model.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(tt.Prob("one"))
}
