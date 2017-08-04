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
		stmtIns, err := dbMaster.Prepare("update t_comico_info set comico_auth=? where comico_id = 77;")
		if err != nil {
			panic(err.Error())
		}
		defer stmtIns.Close()
		res, err := stmtIns.Exec("五仁月饼大大")
		if err != nil {
			panic(err.Error())
		}
		affect, err := res.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		s.ctx.comiservice.logf(LOG_INFO, " affect rows = %d ", affect)
	} else {
		s.ctx.comiservice.logf(LOG_INFO, "dbMaster is nil ", nil)
	}

	channels := "yyyyyyyy"
	fmt.Fprintf(w, "getpost ret = 0")
	return map[string]interface{}{
		"channels": channels,
	}, nil
}
