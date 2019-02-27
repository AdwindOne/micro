/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 2/24/19
 * Time: 23:36
 */
package services

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
