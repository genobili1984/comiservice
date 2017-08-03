package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

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
