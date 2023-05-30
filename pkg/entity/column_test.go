package entity_test

import (
	"testing"

	"github.com/agile-app/flexdb/pkg/entity"
	_ "github.com/agile-app/flexdb/plugins"
	"github.com/funkygao/assert"
)

func TestColumnValidate(t *testing.T) {
	c := &entity.Column{Name: "accountNo", Required: true, Kind: entity.ColumnText}
	assert.Equal(t, nil, c.ValidateCellData("ab"))
	assert.Equal(t, "column[accountNo] is required", c.ValidateCellData("").Error())
}

func TestColumnTableName(t *testing.T) {
	c := &entity.Column{}
	assert.Equal(t, "Field", c.TableName())
}

func TestColumnBeforeCreate(t *testing.T) {
	c := &entity.Column{}
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.Name = " "
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.Name = "x"
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.ModelID = 2
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.Slot = 1
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.Kind = entity.ColumnERP
	assert.Equal(t, false, c.BeforeCreate(nil) != nil)
	c.Name = "a b"
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
	c.Name = "x"
	c.Slot = 1000
	assert.Equal(t, true, c.BeforeCreate(nil) != nil)
}

func TestChoiceOptions(t *testing.T) {
	c := &entity.Column{Choices: "a,", Kind: entity.ColumnChoice}
	assert.Equal(t, 1, len(c.ChoiceOptions()))
	c.Choices = "a,b"
	assert.Equal(t, 2, len(c.ChoiceOptions()))
	c.Choices = ",a"
	assert.Equal(t, 1, len(c.ChoiceOptions()))
}
