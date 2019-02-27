package services

import (
	"net/http"
	"microdemo/micro/accountservice/service"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		service.GetAccount,
	},
}
