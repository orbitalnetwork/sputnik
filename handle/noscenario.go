//+build !scenario

package handle

func ScenarioSupported() bool {
	return false
}

// stubbed
type Scenario struct{}

func (s *Scenario) Init() {
}
