package datetime

import (
	"testing"
	"time"

	"github.com/agile-app/flexdb/internal/spec"
	"github.com/funkygao/assert"
)

func TestParseExcelCell(t *testing.T) {
	cell := "01-06-2021"
	tm, err := time.Parse("01-02-2006", cell)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2021, tm.Year())

	cell = "01-06-21"
	tm, err = time.Parse("01-02-06", cell)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2021, tm.Year())
}

func TestUnixTimestamp(t *testing.T) {
	tm, _ := time.Parse("01-02-06", "01-06-21")
	t.Logf("%v", tm)
	tm1 := time.Unix(1609257600, 0)
	t.Logf("%v %v", tm, tm1)
	t.Logf("%v", tm.Format(spec.YYYYMMDD))

}
