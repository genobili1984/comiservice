package main

import (
	"comiservice/http_api"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	ctx    *Context
	router http.Handler
}

func newHttpServer(ctx *Context) *httpServer {
	log := http_api.Log(ctx.comiservice.logf)

	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	// router.PanicHandler = http_api.LogPanicHandler(ctx.nsqlookupd.logf)
	// router.NotFound = http_api.LogNotFoundHandler(ctx.nsqlookupd.logf)
	// router.MethodNotAllowed = http_api.LogMethodNotAllowedHandler(ctx.nsqlookupd.logf)
	s := &httpServer{
		ctx:    ctx,
		router: router,
	}

	router.Handle("POST", "/epinfo", http_api.Decorate(s.epinfo, log))
	router.Handle("POST", "/getpost", http_api.Decorate(s.getpost, log))
	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}
