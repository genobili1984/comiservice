package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) getpost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	channels := "yyyyyyyy"
	s.ctx.comiservice.logf(LOG_INFO, "channels = %s", channels, nil)
	fmt.Fprintf(w, "getpost ret = 0")
	return map[string]interface{}{
		"channels": channels,
	}, nil
}
