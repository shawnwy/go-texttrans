package texttrans

import (
	"encoding/json"
	"io/ioutil"
)

type modVO struct {
	Matrix         map[string]map[string]float64 `json:"mat"`
	NonPatternProb float64                       `json:"non_pattern_prob"`
	Ngram          int                           `json:"ngram"`
}

func (m modVO) toBigramMatrix() map[int64]float64 {
	var matrix = make(map[int64]float64)
	for kA, runeMat := range m.Matrix {
		a := []rune(kA)[0]
		for kB, prob := range runeMat {
			b := []rune(kB)[0]
			matrix[BigramKey(a, b)] = prob
		}
	}
	return matrix
}

func loadModel(path string) (matrix map[int64]float64, nonPatternProb float64, err error) {
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	var mod modVO
	err = json.Unmarshal(content, &mod)
	if err != nil {
		return
	}
	return mod.toBigramMatrix(), mod.NonPatternProb, nil
}

func BigramKey(a, b rune) int64 {
	return int64(a)<<32 | int64(b)
}
