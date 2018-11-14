//+build !control

package handle

func ControlEnabled() bool {
	return false
}

// stubbed
type Control struct{}

func (c *Control) Init() {
}
