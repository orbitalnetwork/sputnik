package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/coreos/bbolt"
)

type Cache struct {
	boltDB *bolt.DB
}

func (s *Cache) Init(DB string) {
	var err error
	s.boltDB, err = bolt.Open(DB, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Cache) Get() {

	return
}

func (s *Cache) Set() {

	return
}

func (s *Cache) Close() {

	return
}
