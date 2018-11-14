//+build scenario

package handle

import (
	log "github.com/Sirupsen/logrus"
	"github.com/looplab/fsm"
	"github.com/robertkrimen/otto"
)

type Scenario struct {
	js *otto.Otto
}

func (s *Scenario) Init() {
	s.js = otto.New()

	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)
	log.Info("FSM: %#v", fsm)
}

func ScenarioSupported() bool {
	return true
}
