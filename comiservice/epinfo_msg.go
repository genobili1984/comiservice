package main

import (
	"comiservice/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) epinfo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	// reqParams, err := http_api.NewReqParams(req)
	// if err != nil {
	// 	return nil, http_api.Err{400, "INVALID_REQUEST"}
	// }
	// log.Println(reqParams)
	dbMaster := dbmanager.GetDB(dbmanager.DBMaster)
	if dbMaster != nil {
		rows, err := dbMaster.Query("select * from comico_online.t_comico_info limit 5;")
		if err != nil {
			panic(err.Error())
		}
		columns, err := rows.Columns()
		if err != nil {
			panic(err.Error())
		}
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			var value string
			for i, col := range values {
				// Here we can check if the value is nil (NULL value)
				if col == nil {
					value = "NULL"
				} else {
					value = string(col)
				}
				s.ctx.comiservice.logf(LOG_INFO, "%s:%s", columns[i], value, nil)
			}
			fmt.Println("-----------------------------------")
		}
	} else {
		s.ctx.comiservice.logf(LOG_INFO, "dbMaster is nil ", nil)
	}
	channels := "xxxxxxxx"
	return map[string]interface{}{
		"channels": channels,
	}, nil
}
