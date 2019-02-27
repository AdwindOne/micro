package main

import (
	"fmt"
	"microdemo/micro/services"
	"microdemo/micro/accountservice/service"
	"microdemo/micro/accountservice/dbclient"
)

var appName = "accountService"

func main() {
	fmt.Printf("String %v\n",appName)
	initalizeBoltClient()
	services.StartWebServer("6767")
}

func initalizeBoltClient()  {
	service.DBClient = dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
	
}
