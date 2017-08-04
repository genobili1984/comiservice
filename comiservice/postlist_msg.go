package main

import (
	"fmt"
	"net/http"

	"comiservice/db"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) getpost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	dbMaster := dbmanager.GetDB(dbmanager.DBMaster)
	if dbMaster != nil {
		_, err := dbMaster.Exec("update comico_online.t_comico_info set comico_auth = '五仁月饼' where comico_id = 77;")
		if err != nil {
			panic(err.Error())
		}
	} else {
		s.ctx.comiservice.logf(LOG_INFO, "dbMaster is nil ", nil)
	}

	channels := "yyyyyyyy"
	fmt.Fprintf(w, "getpost ret = 0")
	return map[string]interface{}{
		"channels": channels,
	}, nil
}
