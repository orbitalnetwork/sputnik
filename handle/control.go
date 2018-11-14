//+build control

package handle

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
)

type Control struct {
}

func ControlEnabled() bool {
	return true
}

func (c *Control) Init() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

}
