package formula

import (
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/funkygao/assert"
)

func TestGovaluate(t *testing.T) {
	rule := "foo > 10"
	exp, err := govaluate.NewEvaluableExpression(rule) // if rule is 'foo >* 1', will raise err
	assert.Equal(t, nil, err)

	_, err = exp.Eval(nil)
	assert.Equal(t, "No parameter 'foo' found.", err.Error())

	params := map[string]interface{}{
		"foo": 20,
	}
	r, _ := exp.Evaluate(params)
	assert.Equal(t, true, r)
	params["foo"] = 2
	r, _ = exp.Evaluate(params)
	assert.Equal(t, false, r)
}

func BenchmarkGovaluate(b *testing.B) {
	exp := "foo > 10"
	params := map[string]interface{}{
		"foo": 20,
	}
	for i := 0; i < b.N; i++ {
		exp, _ := govaluate.NewEvaluableExpression(exp)
		if true {
			exp.Evaluate(params)
		}
	}
}
