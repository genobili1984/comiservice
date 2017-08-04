package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"comiservice/db"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) getpost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	dbMaster := dbmanager.GetDB(dbmanager.DBMaster)
	if dbMaster != nil {
		timestamps := time.Now().Unix() + (int64)(rand.Intn(1000))
		s.ctx.comiservice.logf(LOG_INFO, " timestamps= %v ", timestamps)
		res, err := dbMaster.Exec("update t_comico_info set comico_auth=? where comico_id = 77;", timestamps)
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
