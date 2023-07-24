# text-trans
text-trans, a simplify go version lib, computes a transition probability of a text which only support ngram=1.

# Usage
## Construct a Texttrans
```golang
import "github.com/shawnwy/go-texttrans"

func main() {
    tt := NewTextTrans("configs/new_model.json")
    prob := tt.Prob("one")
    log.Println("The transition prob of word 'one' is ", prob)
}
```