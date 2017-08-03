package main

import (
	"comiservice/http_api"
	"comiservice/internal"
	"comiservice/util"
	"log"
	"net"
	"os"
	"sync"
)

type ComiService struct {
	sync.RWMutex
	opts         *Options
	httpListener net.Listener
	waitGroup    util.WaitGroupWrapper
}

func New(opts *Options) *ComiService {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, opts.LogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	n := &ComiService{
		opts: opts,
	}

	var err error
	opts.logLevel, err = lg.ParseLogLevel(opts.LogLevel, opts.Verbose)
	if err != nil {
		n.logf(LOG_FATAL, "%s", err)
		os.Exit(1)
	}

	n.logf(LOG_INFO, "comi_service")
	return n
}

func (s *ComiService) Main() {
	ctx := &Context{s}
	httpListener, err := net.Listen("tcp", s.opts.HTTPAddress)
	if err != nil {
		s.logf(LOG_FATAL, "listen (%s) failed - %s", s.opts.HTTPAddress, err)
		os.Exit(1)
	}
	s.Lock()
	s.httpListener = httpListener
	s.Unlock()
	httpServer := newHttpServer(ctx)
	s.waitGroup.Wrap(func() {
		http_api.Serve(httpListener, httpServer, "HTTP", s.logf)
	})
}

func (l *ComiService) Exit() {
	if l.httpListener != nil {
		l.httpListener.Close()
	}
	l.waitGroup.Wait()
}
