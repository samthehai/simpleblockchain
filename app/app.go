package app

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/samthehai/simpleblockchain/server"
)

type app struct {
	servers []server.Server
	errors  chan error
}

func NewApp(servers []server.Server) *app {
	return &app{
		servers: servers,
		errors:  make(chan error, 1+(len(servers)*2)),
	}
}

func (a *app) serve(wg *sync.WaitGroup) {
	for _, s := range a.servers {
		go func(s server.Server) {
			defer wg.Done()
			if err := s.Run(); err != nil {
				a.errors <- err
			}
		}(s)
	}
}

func (a *app) signalHandler(wg *sync.WaitGroup) {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sign := <-ch
		switch sign {
		case syscall.SIGINT, syscall.SIGTERM:
			// this ensures a subsequent INT/TERM will trigger standard go behaviour of
			// terminating.
			signal.Stop(ch)
			a.terminate(wg)
			return
		}
	}
}

func (a *app) terminate(wg *sync.WaitGroup) {
	for _, s := range a.servers {
		go func(s server.Server) {
			defer wg.Done()
			if err := s.Close(); err != nil {
				a.errors <- err
			}
		}(s)
	}
}

func (a *app) Run() error {
	waitdone := make(chan struct{})
	go func() {
		defer close(waitdone)
		var wg sync.WaitGroup
		wg.Add(len(a.servers) * 2)

		go a.signalHandler(&wg)
		a.serve(&wg)

		wg.Wait()
	}()

	select {
	case err := <-a.errors:
		if err == nil {
			panic("unexpected nil error")
		}
		return nil
	case <-waitdone:
		log.Printf("Exiting pid: %v", os.Getegid())
		return nil
	}
}
