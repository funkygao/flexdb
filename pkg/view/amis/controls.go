package amis

func NewDivider() *Control {
	c := NewControl()
	c.Type = "divider"
	return c
}
