package server

import (
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/job"
	"github.com/SantoDE/datahamster/provider"
	"github.com/SantoDE/datahamster/server/http"
	"github.com/SantoDE/datahamster/storage/file"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/store/target"
	"github.com/SantoDE/datahamster/types"
	"time"
)

type Server struct {
	http      *http.Server
	cfg       *configuration.ServerConfiguration
	targets   *types.ConfigTargets
	datastore *store.DataStore
	storage   *file.Storage
}

func NewServer(c configuration.ServerConfiguration, dataStore *store.DataStore) *Server {
	s := new(Server)
	s.http = http.NewServer(c.Address)
	s.cfg = &c
	s.datastore = dataStore
	s.storage = file.NewFileStorage(s.cfg.Storage.File.Path)

	return s
}

func (s *Server) Start(stopChan <-chan bool) {

	var targetsChan = make(chan types.ConfigTargets)

	targetStore := new(target.Store)

	go func() {
		go targetStore.Subscribe(targetsChan)
		if s.cfg.FileTargets != (configuration.FileTargetsConfguration{}) {
			ft := provider.FileTargets{
				s.cfg.FileTargets.File,
				s.cfg.FileTargets.Dir,
			}
			ft.Provide(targetsChan, stopChan)
		}
	}()

	s.http.Init(s.storage, s.datastore, targetStore)

	go s.http.Start()

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		for _, target := range targetStore.ListTargets() {
			if job.ValidateJobNeededByTarget(target, s.datastore) {
				job.Queue(target.Name, s.datastore)
			}
		}
	}
}
