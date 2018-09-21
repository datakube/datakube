package server

import (
	"github.com/datakube/datakube/configuration"
	"github.com/datakube/datakube/job"
	"github.com/datakube/datakube/provider"
	"github.com/datakube/datakube/server/http"
	"github.com/datakube/datakube/storage/file"
	"github.com/datakube/datakube/store"
	"github.com/datakube/datakube/store/target"
	"github.com/datakube/datakube/types"
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

		kp := provider.KubernetesProvider{
			"127.0.0.1:8081",
			"",
			"",
		}
		kp.Provide(targetsChan, stopChan)

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
