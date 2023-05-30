package entity

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestRowSetField(t *testing.T) {
	r := Row{}
	r.SetField(2, "golang")
	assert.Equal(t, "golang", r.S2)
	assert.Equal(t, "golang", r.GetField(2))
	r.SetField(1000, "invalid")
	assert.Equal(t, "", r.GetField(1000))
}

func TestRowGetField(t *testing.T) {
	r := Row{S1: "a", S2: "b"}
	assert.Equal(t, "a", r.GetField(1))
	assert.Equal(t, "b", r.GetField(2))
	assert.Equal(t, "", r.GetField(100))
}

func TestRowBeforeCreate(t *testing.T) {
	r := Row{}
	assert.Equal(t, true, r.MTime.IsZero())
	assert.Equal(t, true, r.CTime.IsZero())

	assert.Equal(t, "empty row.orgID", r.BeforeCreate(nil).Error())
	r.orgID = 2
	assert.Equal(t, "empty row.modelID", r.BeforeCreate(nil).Error())
	r.ModelID = 0
	assert.Equal(t, "empty row.modelID", r.BeforeCreate(nil).Error())
	r.ModelID = 105
	assert.Equal(t, nil, r.BeforeCreate(nil))

	// hook will setup mtime/ctime
	assert.Equal(t, false, r.MTime.IsZero())
	assert.Equal(t, false, r.CTime.IsZero())
	assert.Equal(t, false, r.Deleted)
}

func TestRowBeforeUpdate(t *testing.T) {
}

func TestRowBeforeDelete(t *testing.T) {
}

// 20 ns
func BenchmarkRowSetField(b *testing.B) {
	r := Row{}
	for i := 0; i < b.N; i++ {
		r.SetField(1, "foo")
	}
}

// 18 ns
func BenchmarkRowGetField(b *testing.B) {
	r := Row{S3: "abc"}
	for i := 0; i < b.N; i++ {
		r.GetField(3)
	}
}

func TestRowClob(t *testing.T) {
	r := &RowClob{RowID: 109}
	r.SetField(1, "foo")
	assert.Equal(t, "foo", r.GetField(1))
	assert.Equal(t, "", r.GetField(1999))
	assert.Equal(t, "foo", r.S1)
	r.S1 = "123"
	assert.Equal(t, "123", r.GetField(1))
	r.SetField(2, "bar")
	assert.Equal(t, "bar", r.GetField(2))
}

// 20 ns
func BenchmarkClobSetField(b *testing.B) {
	r := RowClob{}
	for i := 0; i < b.N; i++ {
		r.SetField(1, "foo")
	}
}

// 18 ns
func BenchmarkClobGetField(b *testing.B) {
	r := RowClob{S3: "abc"}
	for i := 0; i < b.N; i++ {
		r.GetField(3)
	}
}
