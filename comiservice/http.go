package main

import (
	"comiservice/http_api"
	"fmt"
	"net/http"

	"time"

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

func (s *httpServer) epinfo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	// reqParams, err := http_api.NewReqParams(req)
	// if err != nil {
	// 	return nil, http_api.Err{400, "INVALID_REQUEST"}
	// }
	// log.Println(reqParams)
	time.Sleep(5e9)
	fmt.Fprintf(w, "epinfo ret = 11111")
	channels := "xxxxxxxx"
	return map[string]interface{}{
		"channels": channels,
	}, nil
}

func (s *httpServer) getpost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	channels := "yyyyyyyy"
	s.ctx.comiservice.logf(LOG_INFO, "channels = %s", channels, nil)
	fmt.Fprintf(w, "getpost ret = 0")
	return map[string]interface{}{
		"channels": channels,
	}, nil
}
