package dto

import (
	"testing"
	"time"

	"github.com/funkygao/assert"
)

func TestBasic(t *testing.T) {
	rd := make(RowData)
	assert.Equal(t, "", rd.StrValueOf("invalid-key"))
	rd.Put("foo", "bar")
	assert.Equal(t, "bar", rd.StrValueOf("foo"))
	assert.Equal(t, false, rd.HasField("invalid_key"))
	assert.Equal(t, true, rd.HasField("foo"))

	assert.Equal(t, uint64(0), rd.ID())
	rd.SetID(2)
	assert.Equal(t, uint64(2), rd.ID())

	t0 := time.Now()
	rd.SetCTime(t0)
	// assert.Equal(t, t0.Unix(), rd.CTime().Unix()) TODO
	rd.SetCUser("cu")
	rd.SetMUser("mu")
	assert.Equal(t, "cu", rd.CUser())
	assert.Equal(t, "mu", rd.MUser())
	assert.Equal(t, true, rd.HasField(KeyCUser))
	t.Logf("%s", rd.JSONString())
}

func BenchmarkRowDataGetString(b *testing.B) {
	rd := make(RowData)
	rd["x"] = "xx"
	for i := 0; i < b.N; i++ {
		rd.StrValueOf("x")
	}
}

func BenchmarkRowDataGetInt(b *testing.B) {
	rd := make(RowData)
	rd["x"] = float64(343)
	for i := 0; i < b.N; i++ {
		rd.StrValueOf("x")
	}
}
