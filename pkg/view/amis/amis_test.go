package amis

import (
	"encoding/json"
	"testing"
)

func TestBasic(t *testing.T) {
	p := NewPage()
	p.Title = "title"
	p.SubTitle = "subtitle"
	p.InitAPI = "localhost"

	p.SetAside("wrapper")
	aside := p.Aside
	aside.Size = "xs"
	nav := NewNav()
	aside.AddBody(nav)
	nav.Stacked = true
	l1 := NewLink("l1", "?id=2")
	nav.AddLink(*l1)
	l2 := NewLink("l2", "?id=3")
	l21 := NewLink("l21", "x")
	l2.AddChild(*l21)
	l22 := NewLink("l22", "y")
	l2.AddChild(*l22)
	nav.AddLink(*l2)

	t1 := &Tpl{Content: "Hello world!", Inline: false}
	p.AddBody(t1)

	b, _ := json.Marshal(p)
	t.Logf("%v", string(b))
}
