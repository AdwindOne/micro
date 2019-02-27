/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 2/25/19
 * Time: 00:05
 */
package service

import (
	"microdemo/micro/accountservice/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

var DBClient dbclient.BoltClient

func GetAccount(w http.ResponseWriter,r *http.Request)  {

	var accountId = mux.Vars(r)["accountId"]

	account, e := DBClient.QueryAccount(accountId)

	if e !=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}