package texttrans

import (
	"math"

	"github.com/shawnwy/go-utils/v5/errors"
)

type TextTrans struct {
	matrix         map[int64]float64
	nonPatternProb float64
}

func NewTextTrans(path string) (*TextTrans, error) {
	matrix, nonPatternProb, err := loadModel(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct TextTrans")
	}
	return &TextTrans{
		matrix:         matrix,
		nonPatternProb: nonPatternProb,
	}, nil
}

func (t *TextTrans) Prob(s string) float64 {
	var logProb float64
	var transCnt float64

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		gramA := runes[i]
		gramB := '\u0000'
		if i+1 != len(runes) {
			gramB = runes[i+1]
		}
		logProb += t.sublineProb(gramA, gramB)
		transCnt += 1
	}
	if transCnt == 0 {
		return 0
	}
	return math.Exp(logProb / transCnt)
}

func (t *TextTrans) sublineProb(a, b rune) float64 {
	if prob, ok := t.matrix[BigramKey(a, b)]; ok {
		return prob
	}
	return t.nonPatternProb
}
